package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/machinebox/graphql"
	"github.com/rohaquinlop/Trainin-exercise/backend/graph/model"
)

func getActualDateTime(value string) string {
	if value == "" {
		date := fmt.Sprintf("%d", time.Now().Unix())
		return date
	} else {
		tmp, _ := time.Parse(time.RFC3339, value)
		date := fmt.Sprintf("%d", tmp.Unix())
		return date
	}
}

func downloadFile(filepath string, url string) error {
	//get file
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//Create file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}

type buyersType []model.Buyer
type productsType []model.Product

func loadBuyers() {
	var data buyersType
	jsonFile, err := ioutil.ReadFile("buyers.json")

	if err != nil {
		log.Printf("Error loading buyers file: %v", err)
	}
	err = json.Unmarshal(jsonFile, &data)
	if err != nil {
		log.Fatal(err)
	}

	client := graphql.NewClient("http://localhost:8080/graphql")

	for _, buyer := range data {
		//make request
		req := graphql.NewRequest(fmt.Sprintf(`
		mutation createBuyer{
			createBuyer(input: {id: "%s", name: "%s", age: %d}){
					name
					age
			}
		}
		`, buyer.ID, buyer.Name, buyer.Age))

		req.Header.Set("Cache-Control", "no-cache")

		ctx := context.Background()
		var respData model.Buyer

		err = client.Run(ctx, req, &respData)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func loadProducts() {
	csvFile, _ := os.Open("products.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = '\''

	client := graphql.NewClient("http://localhost:8080/graphql")

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		req := graphql.NewRequest(fmt.Sprintf(`
		mutation createProduct{
			createProduct(input: {id: "%s", name: "%s", price: %s}){
					name
					price
			}
		}
		`, line[0], line[1], line[2]))

		req.Header.Set("Cache-Control", "no-cache")

		ctx := context.Background()
		var respData model.Buyer

		err = client.Run(ctx, req, &respData)

		if err != nil {
			log.Fatal(err)
		}

	}
}

func main() {
	//Downloading files
	actualDate := getActualDateTime("")

	buyersFile := fmt.Sprintf("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers?date=%s", actualDate)
	productsFile := fmt.Sprintf("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products?date=%s", actualDate)
	transactionsfile := fmt.Sprintf("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions?date=%s", actualDate)

	err := downloadFile("buyers.json", buyersFile)
	if err != nil {
		log.Printf("Error downloading buyers file")
	}
	log.Printf("Downloaded buyers.json")

	err = downloadFile("products.csv", productsFile)
	if err != nil {
		log.Printf("Error downloading products file")
	}
	log.Printf("Downloaded products.csv")

	err = downloadFile("transactions", transactionsfile)
	if err != nil {
		log.Printf("Error downloading transactions file")
	}
	log.Printf("Downloaded transactions")

	//Load data from files
	loadBuyers()
	loadProducts()

}
