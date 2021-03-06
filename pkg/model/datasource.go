package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/grafana/grafana-plugin-model/go/datasource"
)

const (
	AccessMode_LOCAL = "local"
	AccessMode_SFTP  = "sftp"
)

type DatasourceModel struct {
	ID			int64	`json:"id,omitempty"`
	OrgID			int64	`json:"orgId,omitempty"`
	Name			string	`json:"name,omitempty"`
	Type			string	`json:"type:omitempty"`

	Filename		string	`json:"filename"`

	// CSV options
	CsvDelimiter		string	`json:"csvDelimiter"`
	CsvComment		string	`json:"csvComment"`
	CsvTrimLeadingSpace	bool	`json:"csvTrimLeadingSpace"`

	// Access mode: local, sftp
	AccessMode		string	`json:"accessMode"`

	// SFTP
	SftpHost		string	`json:"sftpHost,omitempty"`
	SftpPort		string	`json:"sftpPort,omitempty"`
	SftpUser		string	`json:"sftUser,omitempty"`
	SftpPassword		string	`json:"sftPassword,omitempty"`
	SftpWorkingDir		string	`json:"sftpWorkingDir"`		// Local working dir
	SftpIgnoreHostKey	bool	`json:"sftpIgnoreHostKey"`
}

func CreateDatasourceFrom(req datasource.DatasourceRequest) (*DatasourceModel, error) {
	model := &DatasourceModel{}
	err := json.Unmarshal([]byte(req.Datasource.JsonData), &model)
	if err != nil {
		return nil, err
	}

	model.ID = req.Datasource.Id
	model.OrgID = req.Datasource.OrgId
	model.Name = req.Datasource.Name
	model.Type = req.Datasource.Type
	model.SftpPassword = req.Datasource.DecryptedSecureJsonData["sftpPassword"]

	if len(model.CsvDelimiter) == 0 {
		model.CsvDelimiter = ","
	}

	if len(model.CsvComment) == 0 {
		model.CsvComment = "#"
	}

	if model.AccessMode == AccessMode_SFTP {
		if len(model.SftpPort) == 0 {
			model.SftpPort = "22"
		}
	}

  	return model, validateDatasourceModel(model)
}

func (m *DatasourceModel) String() string {
	jsonBytes, _ := json.Marshal(m)
	return string(jsonBytes)
}

func validateDatasourceModel(model *DatasourceModel) error {
	if model.AccessMode == AccessMode_LOCAL {
		return validateLocalDatasourceModel(model)
	}
	if model.AccessMode == AccessMode_SFTP {
		return validateSftpDatasourceModel(model)
	}
	return errors.New(fmt.Sprintf("unknown access mode `%s`", model.AccessMode))
}

func validateLocalDatasourceModel(model *DatasourceModel) error {
	if len(model.Filename) == 0 {
		return errors.New("the path to the CSV file is not defined")
	}
	return nil
}

func validateSftpDatasourceModel(model *DatasourceModel) error {
	if len(model.Filename) == 0 {
		return errors.New("the path to the CSV file is not defined")
	}
	if len(model.SftpHost) == 0 {
		return errors.New("SFTP hot is not defined")
	}
	if len(model.SftpWorkingDir) == 0 {
		return errors.New("working dir is not defined")
	}
	return nil
}
