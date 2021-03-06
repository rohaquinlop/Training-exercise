package loaddata

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
	"os/exec"
	"strings"
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
		log.Printf("Error loading buyers json file: %v", err)
	}

	client := graphql.NewClient("http://localhost:5656/graphql")

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

	}
}

func loadProducts() {
	csvFile, _ := os.Open("products.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = '\''

	client := graphql.NewClient("http://localhost:5656/graphql")

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

	}
}

func loadTransactions() {
	_, err := exec.Command("/bin/sh", "./loaddata/makeStandard.sh").Output()

	if err != nil {
		log.Fatal(err)
	}

	csvFile, _ := os.Open("transactions.csv")

	reader := csv.NewReader(bufio.NewReader(csvFile))

	reader.Comma = ';'
	client := graphql.NewClient("http://localhost:5656/graphql")
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		line[4] = strings.Replace(line[4], "(", "", -1)
		line[4] = strings.Replace(line[4], ")", "", -1)

		tmp := strings.Split(line[4], ",")
		productIds := "["

		for i, productId := range tmp {
			productIds = productIds + fmt.Sprintf("\"%s\"", productId)
			if i < len(tmp)-1 {
				productIds = productIds + ","
			}
		}
		productIds = productIds + "]"

		req := graphql.NewRequest(fmt.Sprintf(`
		mutation createTransaction{
			createTransaction(input: {id: "%s", buyerId: "%s", ip: "%s", device: "%s", productIds: %s}){
					id
					ip
			}
		}
		`, line[0], line[1], line[2], line[3], productIds))

		req.Header.Set("Cache-Control", "no-cache")

		ctx := context.Background()
		var respData model.Buyer

		err = client.Run(ctx, req, &respData)

	}
}

func LoadDataDB(date string) {
	//Downloading files
	actualDate := getActualDateTime(date)

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
	log.Printf("Buyers loaded successfully")
	loadProducts()
	log.Printf("Products loaded successfully")
	loadTransactions()
	log.Printf("Transactions loaded successfully")

}
