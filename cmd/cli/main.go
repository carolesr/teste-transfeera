package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/teste-transfeera/internal/entity"
	"github.com/teste-transfeera/internal/model"
	"github.com/teste-transfeera/pkg/shared"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	rootCmd := commands["root"]
	rootCmd.AddCommand(commands["seed"])
	err = rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

var commands = map[string]*cobra.Command{
	"root": {
		Use:   "root",
		Short: "A CLI to execute basic configuration for the API",
	},
	"seed": {
		Use:   "seed",
		Short: "Seeds the database with 30 records of Receiver",
		Run:   seed,
	},
}

func seed(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(os.Getenv("DATABASE_URL"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("transfeera").Collection("receiver")

	err = db.Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Dropped Receiver collection")

	receiversToInsert := receivers()

	_, err = db.InsertMany(ctx, receiversToInsert)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Receivers inserted successfully!")
}

func receivers() []interface{} {
	return []interface{}{
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "290.551.590-26",
			Name:       "Receiver 1",
			Email:      "RECEIVER1@GMAIL.COM",
			Bank:       shared.GetPointerStr("Bradesco"),
			Agency:     shared.GetPointerStr("0814-0"),
			Account:    shared.GetPointerStr("01002713-9"),
			Status:     string(entity.Draft),
			Pix: model.Pix{
				KeyType: string(entity.CPF),
				Key:     "290.551.590-26",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "516.488.970-61",
			Name:       "Receiver 2",
			Email:      "RECEIVER2@GMAIL.COM",
			Bank:       shared.GetPointerStr("Banco do Brasil"),
			Agency:     shared.GetPointerStr("8016"),
			Account:    shared.GetPointerStr("1051790-1"),
			Status:     string(entity.Validated),
			Pix: model.Pix{
				KeyType: string(entity.CPF),
				Key:     "516.488.970-61",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "300.227.450-09",
			Name:       "Receiver 3",
			Email:      "RECEIVER3@GMAIL.COM",
			Bank:       shared.GetPointerStr("Bradesco"),
			Agency:     shared.GetPointerStr("3073"),
			Account:    shared.GetPointerStr("0771847-0"),
			Status:     string(entity.Draft),
			Pix: model.Pix{
				KeyType: string(entity.CPF),
				Key:     "300.227.450-09",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "395.354.370-97",
			Name:       "Receiver 4",
			Email:      "RECEIVER4@GMAIL.COM",
			Bank:       shared.GetPointerStr("Bradesco"),
			Agency:     shared.GetPointerStr("2215"),
			Account:    shared.GetPointerStr("1677453-7"),
			Status:     string(entity.Validated),
			Pix: model.Pix{
				KeyType: string(entity.CPF),
				Key:     "395.354.370-97",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "750.941.900-08",
			Name:       "Receiver 5",
			Email:      "RECEIVER5@GMAIL.COM",
			Bank:       shared.GetPointerStr("Santander"),
			Agency:     shared.GetPointerStr("0485"),
			Account:    shared.GetPointerStr("53311681-7"),
			Status:     string(entity.Draft),
			Pix: model.Pix{
				KeyType: string(entity.CPF),
				Key:     "750.941.900-08",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "454.859.820-00",
			Name:       "Receiver 6",
			Email:      "RECEIVER6@GMAIL.COM",
			Bank:       shared.GetPointerStr("Banco do Brasil"),
			Agency:     shared.GetPointerStr("0814-5"),
			Account:    shared.GetPointerStr("5448"),
			Status:     string(entity.Validated),
			Pix: model.Pix{
				KeyType: string(entity.CPF),
				Key:     "454.859.820-00",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "40.424.263/0001-42",
			Name:       "Receiver 7",
			Email:      "RECEIVER7@GMAIL.COM",
			Bank:       shared.GetPointerStr("Bradesco"),
			Agency:     shared.GetPointerStr("1674"),
			Account:    shared.GetPointerStr("0722375-7"),
			Status:     string(entity.Draft),
			Pix: model.Pix{
				KeyType: string(entity.CNPJ),
				Key:     "40.424.263/0001-42",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "45.325.641/0001-54",
			Name:       "Receiver 8",
			Email:      "RECEIVER8@GMAIL.COM",
			Bank:       shared.GetPointerStr("Bradesco"),
			Agency:     shared.GetPointerStr("1515"),
			Account:    shared.GetPointerStr("1858481-6"),
			Status:     string(entity.Validated),
			Pix: model.Pix{
				KeyType: string(entity.CNPJ),
				Key:     "45.325.641/0001-54",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "08.219.094/0001-04",
			Name:       "Receiver 9",
			Email:      "RECEIVER9@GMAIL.COM",
			Bank:       shared.GetPointerStr("Banco do Brasil"),
			Agency:     shared.GetPointerStr("8016"),
			Account:    shared.GetPointerStr("298417-2"),
			Status:     string(entity.Draft),
			Pix: model.Pix{
				KeyType: string(entity.CNPJ),
				Key:     "08.219.094/0001-04",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "60.686.639/0001-02",
			Name:       "Receiver 10",
			Email:      "RECEIVER10@GMAIL.COM",
			Bank:       shared.GetPointerStr("Itaú"),
			Agency:     shared.GetPointerStr("5586"),
			Account:    shared.GetPointerStr("49718-1"),
			Status:     string(entity.Validated),
			Pix: model.Pix{
				KeyType: string(entity.CNPJ),
				Key:     "60.686.639/0001-02",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "14.890.924/0001-15",
			Name:       "Receiver 11",
			Email:      "RECEIVER11@GMAIL.COM",
			Bank:       shared.GetPointerStr("Banco do Brasil"),
			Agency:     shared.GetPointerStr("3320"),
			Account:    shared.GetPointerStr("1179294-9"),
			Status:     string(entity.Draft),
			Pix: model.Pix{
				KeyType: string(entity.CNPJ),
				Key:     "14.890.924/0001-15",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "38.325.271/0001-90",
			Name:       "Receiver 12",
			Email:      "RECEIVER12@GMAIL.COM",
			Bank:       shared.GetPointerStr("Santander"),
			Agency:     shared.GetPointerStr("0947"),
			Account:    shared.GetPointerStr("43866736-7"),
			Status:     string(entity.Validated),
			Pix: model.Pix{
				KeyType: string(entity.CNPJ),
				Key:     "38.325.271/0001-90",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "800.686.200-12",
			Name:       "Receiver 13",
			Email:      "RECEIVER13@GMAIL.COM",
			Bank:       shared.GetPointerStr("Banco do Brasil"),
			Agency:     shared.GetPointerStr("1404"),
			Account:    shared.GetPointerStr("1218287-7"),
			Status:     string(entity.Draft),
			Pix: model.Pix{
				KeyType: string(entity.Email),
				Key:     "RECEIVER13@GMAIL.COM",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "586.076.790-07",
			Name:       "Receiver 14",
			Email:      "RECEIVER14@GMAIL.COM",
			Bank:       shared.GetPointerStr("Santander"),
			Agency:     shared.GetPointerStr("1728"),
			Account:    shared.GetPointerStr("27645921-0"),
			Status:     string(entity.Validated),
			Pix: model.Pix{
				KeyType: string(entity.Email),
				Key:     "RECEIVER14@GMAIL.COM",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "259.498.450-72",
			Name:       "Receiver 15",
			Email:      "RECEIVER15@GMAIL.COM",
			Bank:       shared.GetPointerStr("Santander"),
			Agency:     shared.GetPointerStr("2210"),
			Account:    shared.GetPointerStr("35155013-6"),
			Status:     string(entity.Draft),
			Pix: model.Pix{
				KeyType: string(entity.Email),
				Key:     "RECEIVER15@GMAIL.COM",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "861.248.030-20",
			Name:       "Receiver 16",
			Email:      "RECEIVER16@GMAIL.COM",
			Bank:       shared.GetPointerStr("Santander"),
			Agency:     shared.GetPointerStr("1194"),
			Account:    shared.GetPointerStr("46976438-8"),
			Status:     string(entity.Validated),
			Pix: model.Pix{
				KeyType: string(entity.Email),
				Key:     "RECEIVER16@GMAIL.COM",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "919.502.190-62",
			Name:       "Receiver 17",
			Email:      "RECEIVER17@GMAIL.COM",
			Bank:       shared.GetPointerStr("Santander"),
			Agency:     shared.GetPointerStr("3731"),
			Account:    shared.GetPointerStr("60764032-4"),
			Status:     string(entity.Draft),
			Pix: model.Pix{
				KeyType: string(entity.Email),
				Key:     "RECEIVER17@GMAIL.COM",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "952.497.300-60",
			Name:       "Receiver 18",
			Email:      "RECEIVER18@GMAIL.COM",
			Bank:       shared.GetPointerStr("Bradesco"),
			Agency:     shared.GetPointerStr("2961"),
			Account:    shared.GetPointerStr("1276583-5"),
			Status:     string(entity.Validated),
			Pix: model.Pix{
				KeyType: string(entity.Email),
				Key:     "RECEIVER18@GMAIL.COM",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "84.181.527/0001-50",
			Name:       "Receiver 19",
			Email:      "RECEIVER19@GMAIL.COM",
			Bank:       shared.GetPointerStr("Banco do Brasil"),
			Agency:     shared.GetPointerStr("2750"),
			Account:    shared.GetPointerStr("122810-2"),
			Status:     string(entity.Draft),
			Pix: model.Pix{
				KeyType: string(entity.Phone),
				Key:     "5548",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "65.197.494/0001-91",
			Name:       "Receiver 20",
			Email:      "RECEIVER20@GMAIL.COM",
			Bank:       shared.GetPointerStr("Bradesco"),
			Agency:     shared.GetPointerStr("0606"),
			Account:    shared.GetPointerStr("0436294-2"),
			Status:     string(entity.Validated),
			Pix: model.Pix{
				KeyType: string(entity.Phone),
				Key:     "5548",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "29.516.384/0001-81",
			Name:       "Receiver 21",
			Email:      "RECEIVER21@GMAIL.COM",
			Bank:       shared.GetPointerStr("Santander"),
			Agency:     shared.GetPointerStr("0500"),
			Account:    shared.GetPointerStr("50585125-8"),
			Status:     string(entity.Draft),
			Pix: model.Pix{
				KeyType: string(entity.Phone),
				Key:     "5548",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "24.269.544/0001-11",
			Name:       "Receiver 22",
			Email:      "RECEIVER22@GMAIL.COM",
			Bank:       shared.GetPointerStr("Itaú"),
			Agency:     shared.GetPointerStr("0289"),
			Account:    shared.GetPointerStr("0606476-0"),
			Status:     string(entity.Validated),
			Pix: model.Pix{
				KeyType: string(entity.Phone),
				Key:     "5548",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "64.004.460/0001-70",
			Name:       "Receiver 23",
			Email:      "RECEIVER23@GMAIL.COM",
			Bank:       shared.GetPointerStr("Itaú"),
			Agency:     shared.GetPointerStr("9688"),
			Account:    shared.GetPointerStr("83438-2"),
			Status:     string(entity.Draft),
			Pix: model.Pix{
				KeyType: string(entity.Phone),
				Key:     "5548",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "77.334.798/0001-32",
			Name:       "Receiver 24",
			Email:      "RECEIVER24@GMAIL.COM",
			Bank:       shared.GetPointerStr("Bradesco"),
			Agency:     shared.GetPointerStr("3522"),
			Account:    shared.GetPointerStr("0507968-3"),
			Status:     string(entity.Validated),
			Pix: model.Pix{
				KeyType: string(entity.Phone),
				Key:     "5548",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "273.753.420-83",
			Name:       "Receiver 25",
			Email:      "RECEIVER25@GMAIL.COM",
			Bank:       shared.GetPointerStr("Santander"),
			Agency:     shared.GetPointerStr("2030"),
			Account:    shared.GetPointerStr("48638554-6"),
			Status:     string(entity.Draft),
			Pix: model.Pix{
				KeyType: string(entity.RandomKey),
				Key:     "41fcd1dc-ccf5-5ef3-97b8-6254ed1f5dbf",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "300.258.870-92",
			Name:       "Receiver 26",
			Email:      "RECEIVER26@GMAIL.COM",
			Bank:       shared.GetPointerStr("Banco do Brasil"),
			Agency:     shared.GetPointerStr("0732"),
			Account:    shared.GetPointerStr("1266018-3"),
			Status:     string(entity.Validated),
			Pix: model.Pix{
				KeyType: string(entity.RandomKey),
				Key:     "41fcd1dc-ccf5-5ef3-97b8-6254ed1f5dbf",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "142.338.070-32",
			Name:       "Receiver 27",
			Email:      "RECEIVER27@GMAIL.COM",
			Bank:       shared.GetPointerStr("Bradesco"),
			Agency:     shared.GetPointerStr("3376"),
			Account:    shared.GetPointerStr("0128532-7"),
			Status:     string(entity.Draft),
			Pix: model.Pix{
				KeyType: string(entity.RandomKey),
				Key:     "41fcd1dc-ccf5-5ef3-97b8-6254ed1f5dbf",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "770.656.270-04",
			Name:       "Receiver 28",
			Email:      "RECEIVER28@GMAIL.COM",
			Bank:       shared.GetPointerStr("Santander"),
			Agency:     shared.GetPointerStr("3332"),
			Account:    shared.GetPointerStr("37155495-0"),
			Status:     string(entity.Validated),
			Pix: model.Pix{
				KeyType: string(entity.RandomKey),
				Key:     "41fcd1dc-ccf5-5ef3-97b8-6254ed1f5dbf",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "760.572.110-22",
			Name:       "Receiver 29",
			Email:      "RECEIVER29@GMAIL.COM",
			Bank:       shared.GetPointerStr("Bradesco"),
			Agency:     shared.GetPointerStr("6158"),
			Account:    shared.GetPointerStr("0107178-5"),
			Status:     string(entity.Draft),
			Pix: model.Pix{
				KeyType: string(entity.RandomKey),
				Key:     "41fcd1dc-ccf5-5ef3-97b8-6254ed1f5dbf",
			},
			CreatedAt: time.Now(),
		},
		model.Receiver{
			ID:         primitive.NewObjectID(),
			Identifier: "788.253.700-40",
			Name:       "Receiver 30",
			Email:      "RECEIVER30@GMAIL.COM",
			Bank:       shared.GetPointerStr("Banco do Brasil"),
			Agency:     shared.GetPointerStr("4529"),
			Account:    shared.GetPointerStr("54114-1"),
			Status:     string(entity.Validated),
			Pix: model.Pix{
				KeyType: string(entity.RandomKey),
				Key:     "41fcd1dc-ccf5-5ef3-97b8-6254ed1f5dbf",
			},
			CreatedAt: time.Now(),
		},
	}
}
