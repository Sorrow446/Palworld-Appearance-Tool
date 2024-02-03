package main

import (
	"compress/zlib"
	"fmt"
	"os/exec"
	"bytes"
	"os"
	"io"
	"encoding/json"
	"errors"
	"strings"
	"encoding/binary"

	"github.com/alexflint/go-arg"
)

func readSaveData(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	_, err = f.Seek(12, io.SeekCurrent)
	if err != nil {
		return nil, err
	}
	r, err := zlib.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	decompData, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return decompData, nil
}

func ueSaveImport(decompData []byte) ([]byte, error) {
	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	cmd := exec.Command("./uesave.exe", "from-json")
	cmd.Stdin = bytes.NewReader(decompData)
	cmd.Stderr = &errBuffer
	cmd.Stdout = &outBuffer
	err := cmd.Run()
	if err != nil {
		errString := fmt.Sprintf("%s\n%s", err, errBuffer.String())
		return nil, errors.New(errString)
	}
	return outBuffer.Bytes(), nil
}

func ueSaveExport(decompData []byte) ([]byte, error) {
	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	cmd := exec.Command("./uesave.exe", "to-json")
	cmd.Stdin = bytes.NewReader(decompData)
	cmd.Stderr = &errBuffer
	cmd.Stdout = &outBuffer
	err := cmd.Run()
	if err != nil {
		errString := fmt.Sprintf("%s\n%s", err, errBuffer.String())
		return nil, errors.New(errString)
	}
	return outBuffer.Bytes(), nil
}

func getAppDataMap(m map[string]any) map[string]any {
	m = m["root"].(map[string]any)["properties"].(map[string]any)["SaveData"].
		(map[string]any)["Struct"].(map[string]any)["value"].
		(map[string]any)["Struct"].(map[string]any)["PlayerCharacterMakeData"].
		(map[string]any)["Struct"].(map[string]any)["value"].
		(map[string]any)["Struct"].(map[string]any)
	return m
}

func parseAppData(data []byte) (map[string]any, *CharAppearanceData, error) {
	var fullMap map[string]any
	err := json.Unmarshal(data, &fullMap)
	if err != nil {
		return nil, nil, err
	}

	appDataMap := getAppDataMap(fullMap)

	appData, err := json.Marshal(appDataMap)
	if err != nil {
		return nil, nil, err
	}

	var appDataObj CharAppearanceData
	err = json.Unmarshal(appData, &appDataObj)
	if err != nil {
		return nil, nil, err
	}
	return fullMap, &appDataObj, nil
}

func setNewAppData(m map[string]any, appData *CharAppearanceData) map[string]any {
	m["root"].(map[string]any)["properties"].(map[string]any)["SaveData"].
		(map[string]any)["Struct"].(map[string]any)["value"].
		(map[string]any)["Struct"].(map[string]any)["PlayerCharacterMakeData"].
		(map[string]any)["Struct"].(map[string]any)["value"].
		(map[string]any)["Struct"] = appData
	return m	
}

func writeToJSON(outPath string, appData *CharAppearanceData) error {
	data, err := json.MarshalIndent(appData, "", "\t")
	if err != nil {
		return err
	}
	err = os.WriteFile(outPath, data, 0666)
	return err
}

func parseArgs() *Args {
	var args Args
	arg.MustParse(&args)
	args.Command = strings.ToLower(args.Command)
	return &args
}

func exportToJSON(args *Args) error {

	ok := strings.HasSuffix(args.InPath, ".sav") && strings.HasSuffix(args.OutPath, ".json")
	if !ok {
		return errors.New(
			"invalid file extension pair, expected .sav for in and .json for out",
		)
	}

	decompData, err := readSaveData(args.InPath)
	if err != nil {
		return err
	}

	parsedGSAVData, err := ueSaveExport(decompData)
	if err != nil {
		return err
	}

	_, appData, err := parseAppData(parsedGSAVData)
	if err != nil {
		return err
	}
	err = writeToJSON(args.OutPath, appData)
	return err
}

func readAppDataJSON(path string) (*CharAppearanceData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var obj CharAppearanceData
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func putI32LE(n int) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(n))
	return buf
}

func writeNewSave(path string, m map[string]any) error {
    data, err := json.Marshal(m)
    if err != nil {
        return err
    }

    data2, err := ueSaveImport(data)
    if err != nil {
    	return err
    }   

    f, err := os.Create(path)
	if err != nil {
		return err
    }
    defer f.Close()   

    uncompSize := putI32LE(len(data2))

    _, err = f.Write(uncompSize)
    if err != nil {
        return err
    }

    var compData bytes.Buffer

    w := zlib.NewWriter(&compData)
    _, err = w.Write(data2)
    w.Close()
    if err != nil {
       return err
    }


    compSize := putI32LE(compData.Len())

    _, err = f.Write(compSize)
    if err != nil {
        return err
    }

    _, err = f.WriteString("PlZ1")
    if err != nil {
        return err
    }

    _, err = f.Write(compData.Bytes())
    if err != nil {
        return err
    }   
    return err
}

func importToSave(args *Args) error {
	ok := strings.HasSuffix(args.InPath, ".json") && strings.HasSuffix(args.OutPath, ".sav")
	if !ok {
		return errors.New(
			"invalid file extension pair, expected .json for in and .sav for out",
		)
	}
	decompData, err := readSaveData(args.OutPath)
	if err != nil {
		return err
	}

	parsedGSAVData, err := ueSaveExport(decompData)
	if err != nil {
		return err
	}

	m, _, err := parseAppData(parsedGSAVData)
	if err != nil {
		return err
	}
	newAppData, err := readAppDataJSON(args.InPath)
	if err != nil {
		return err
	}
	m = setNewAppData(m, newAppData)
	err = writeNewSave(args.OutPath, m)
	return err
}

func main() {
	args := parseArgs()

	var err error

	switch args.Command {
	case "import":
		err = importToSave(args)
	case "export":
		err = exportToJSON(args)
	default:
		panic("unknown command: " + args.Command)
	}
	if err != nil {
		panic(err)
	}
}