package services

import "simple-microservice/internal/model"

var Accounts []model.AccountInfo

func init() {
	Accounts = []model.AccountInfo{
		{
			Name:      "Prateek",
			CRN:       12323532,
			Balance:   10000,
			Holdings:  5000,
			AccType:   model.SAVING,
			AccHealth: model.OK,
		},
		{
			Name:      "Mandal",
			CRN:       564678345,
			Balance:   10000,
			Holdings:  5000,
			AccType:   model.SAVING,
			AccHealth: model.OK,
		},
		{
			Name:      "Poulami",
			CRN:       14654758353,
			Balance:   10000,
			Holdings:  5000,
			AccType:   model.SAVING,
			AccHealth: model.OK,
		},
		{
			Name:      "Manna",
			CRN:       253454364,
			Balance:   10000,
			Holdings:  5000,
			AccType:   model.SAVING,
			AccHealth: model.OK,
		},
		{
			Name:      "Steve",
			CRN:       234563462,
			Balance:   10000,
			Holdings:  5000,
			AccType:   model.GL,
			AccHealth: model.WARN,
		},
		{
			Name:      "Gerard",
			CRN:       68765545,
			Balance:   10000,
			Holdings:  5000,
			AccType:   model.GL,
			AccHealth: model.OK,
		},
		{
			Name:      "Butler",
			CRN:       54675673,
			Balance:   10000,
			Holdings:  5000,
			AccType:   model.CURRENT,
			AccHealth: model.OK,
		},
		{
			Name:      "Monka",
			CRN:       86786564,
			Balance:   10000,
			Holdings:  5000,
			AccType:   model.OVERDRAFT,
			AccHealth: model.SUB1,
		},
		{
			Name:      "Monke",
			CRN:       36747353,
			Balance:   10000,
			Holdings:  5000,
			AccType:   model.OVERDRAFT,
			AccHealth: model.DELINQUENT,
		},
		{
			Name:      "Casandra",
			CRN:       36547325,
			Balance:   10000,
			Holdings:  5000,
			AccType:   model.OVERDRAFT,
			AccHealth: model.OVERDUE30,
		},
		{
			Name:      "Mongo",
			CRN:       35465735,
			Balance:   10000,
			Holdings:  5000,
			AccType:   model.OVERDRAFT,
			AccHealth: model.WARN,
		},
		{
			Name:      "Redis",
			CRN:       246375378,
			Balance:   10000,
			Holdings:  5000,
			AccType:   model.OVERDRAFT,
			AccHealth: model.SUB3,
		},
	}
}
