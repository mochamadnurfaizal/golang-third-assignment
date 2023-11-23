package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang-third-assignment/config"
	"golang-third-assignment/models"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UpdateEnvirontmentByGorm(ctx echo.Context) error {

	db := config.GetDB()

	id := ctx.Param("id")
	if id == "" {
		return GenerateErrorResponse(ctx, "Data ID Tidak Ditemukan")
	}

	paramID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return GenerateErrorResponse(ctx, err.Error())
	}

	Environtment := models.Environtment{}

	if err := ctx.Bind(&Environtment); err != nil {
		fmt.Println(err)
		return GenerateErrorResponse(ctx, "Format Data Salah")
	}

	Environtment.ID = uint(paramID)

	err = db.Save(&Environtment).Error
	if err != nil {
		fmt.Println(err)
		return GenerateErrorResponse(ctx, err.Error())
	}

	return GenerateSuccessResponse(ctx, "Update Data Success", Environtment)
}

var PORT = ":8088"
var POSTURL = "http://localhost" + PORT + "/updateenvirontment/" + ID_ENVIRONTMENT
var METHOD = "POST"
var ID_ENVIRONTMENT = "1"

func UpdateData() {

	body, err := json.Marshal(map[string]int{
		"water": rand.Intn(100-1) + 1,
		"wind":  rand.Intn(100-1) + 1,
	})

	if err != nil {
		panic(err)
	}

	r, err := http.NewRequest(METHOD, POSTURL, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	if res.Body == nil {
		panic(errors.New("Gagal Mendapatkan Body"))
	}

	if res.StatusCode != http.StatusOK {
		panic(res.Status)
	}

	printEnvirontment(res)

	defer res.Body.Close()
}

func printEnvirontment(res *http.Response) {
	var response models.Response
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		panic(err)
	}

	environtment := getEnvirontmentFromDTO(response.Data)

	responseByte, err := json.Marshal(map[string]int{
		"water": environtment.Water,
		"wind":  environtment.Wind,
	})

	if err != nil {
		panic(err)
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, responseByte, "", "\t")
	if err != nil {
		panic(err)
	}

	statusWater, statusWind := cekEnvirontment(environtment)

	fmt.Println(prettyJSON.String())
	fmt.Println("status water :", statusWater)
	fmt.Println("status wind :", statusWind)
	fmt.Println()
}

func getEnvirontmentFromDTO(data interface{}) models.Environtment {
	m := data.(map[string]interface{})

	water := m["water"].(float64)
	wind := m["wind"].(float64)

	environtment := models.Environtment{
		Water: int(water),
		Wind:  int(wind),
	}

	return environtment
}

func cekEnvirontment(environtment models.Environtment) (string, string) {
	var statusWater string
	var statusWind string

	switch {
	case environtment.Water <= 5:
		statusWater = "aman"
	case environtment.Water >= 6 && environtment.Water <= 8:
		statusWater = "siaga"
	case environtment.Water > 8:
		statusWater = "bahaya"
	default:
		statusWater = "data tidak ditemukan"
	}

	switch {
	case environtment.Wind <= 6:
		statusWind = "aman"
	case environtment.Wind >= 7 && environtment.Wind <= 15:
		statusWind = "siaga"
	case environtment.Wind > 15:
		statusWind = "bahaya"
	default:
		statusWind = "data tidak ditemukan"
	}

	return statusWater, statusWind

}
