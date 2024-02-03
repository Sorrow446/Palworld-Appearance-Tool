package main

type Args struct {
	Command string `arg:"positional, required" help:"import/export\n\t\t\t Import JSON appearance data to save file or export JSON appearance data from save file."`
	InPath  string `arg:"-i" help:"Input path."`
	OutPath string `arg:"-o" help:"Output path."`
}

type CharAppearanceData struct {
	BodyMeshName struct {
		Name struct {
			Value string `json:"value"`
		} `json:"Name"`
	} `json:"BodyMeshName"`
	HeadMeshName struct {
		Name struct {
			Value string `json:"value"`
		} `json:"Name"`
	} `json:"HeadMeshName"`
	HairMeshName struct {
		Name struct {
			Value string `json:"value"`
		} `json:"Name"`
	} `json:"HairMeshName"`
	HairColor struct {
		Struct struct {
			Value struct {
				LinearColor struct {
					R float64 `json:"r"`
					G float64 `json:"g"`
					B float64 `json:"b"`
					A float64 `json:"a"`
				} `json:"LinearColor"`
			} `json:"value"`
			StructType string `json:"struct_type"`
			StructID   string `json:"struct_id"`
		} `json:"Struct"`
	} `json:"HairColor"`
	BrowColor struct {
		Struct struct {
			Value struct {
				LinearColor struct {
					R float64 `json:"r"`
					G float64 `json:"g"`
					B float64 `json:"b"`
					A float64 `json:"a"`
				} `json:"LinearColor"`
			} `json:"value"`
			StructType string `json:"struct_type"`
			StructID   string `json:"struct_id"`
		} `json:"Struct"`
	} `json:"BrowColor"`
	BodyColor struct {
		Struct struct {
			Value struct {
				LinearColor struct {
					R float64 `json:"r"`
					G float64 `json:"g"`
					B float64 `json:"b"`
					A float64 `json:"a"`
				} `json:"LinearColor"`
			} `json:"value"`
			StructType string `json:"struct_type"`
			StructID   string `json:"struct_id"`
		} `json:"Struct"`
	} `json:"BodyColor"`
	BodySubsurfaceColor struct {
		Struct struct {
			Value struct {
				LinearColor struct {
					R float64 `json:"r"`
					G float64 `json:"g"`
					B float64 `json:"b"`
					A float64 `json:"a"`
				} `json:"LinearColor"`
			} `json:"value"`
			StructType string `json:"struct_type"`
			StructID   string `json:"struct_id"`
		} `json:"Struct"`
	} `json:"BodySubsurfaceColor"`
	EyeColor struct {
		Struct struct {
			Value struct {
				LinearColor struct {
					R float64 `json:"r"`
					G float64 `json:"g"`
					B float64 `json:"b"`
					A float64 `json:"a"`
				} `json:"LinearColor"`
			} `json:"value"`
			StructType string `json:"struct_type"`
			StructID   string `json:"struct_id"`
		} `json:"Struct"`
	} `json:"EyeColor"`
	EyeMaterialName struct {
		Name struct {
			Value string `json:"value"`
		} `json:"Name"`
	} `json:"EyeMaterialName"`
	VoiceID struct {
		Int struct {
			Value int `json:"value"`
		} `json:"Int"`
	} `json:"VoiceID"`
}