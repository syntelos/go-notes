/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package notes

type Table string
type TableName Table
type TablePath Table
type TableLink Table

const (
	TableNamePolitics	TableName	= "politics"
	TableNameSociology	TableName	= "sociology"
	TableNameTheTempleOfAthena	TableName	= "the_temple_of_athena"

	TablePathPolitics	TablePath	= "/gegonen /politics"
	TablePathSociology	TablePath	= "/syntelos /science /anthropology /sociology"
	TablePathTheTempleOfAthena	TablePath	= "/gegonen /the_temple_of_athena"

	TableLinkPolitics	TableLink	= "https://drive.google.com/drive/folders/1uhgjgL8HBzRwCgP6pXpCwQqLZ0IC_pgJ?usp=drive_link"
	TableLinkSociology	TableLink	= "https://drive.google.com/drive/folders/1etCIitYhVQH8_Wf6oHeu_rWUmDfTYZD-"
	TableLinkTheTempleOfAthena	TableLink	= "https://drive.google.com/drive/folders/1nUpMgy9n-wHWN13hlXm-YZH5rojTUi0j"
)

func IsTableName(name TableName) bool {
	switch name {
	case TableNamePolitics, TableNameSociology, TableNameTheTempleOfAthena:
		return true
	default:
		return false
	}
}

func (this TableName) Path() TablePath {
	switch this {
	case TableNamePolitics:
		return TablePathPolitics
	case TableNameSociology:
		return TablePathSociology
	case TableNameTheTempleOfAthena:
		return TablePathTheTempleOfAthena
	default:
		return ""
	}
}

func (this TableName) Link() TableLink {
	switch this {
	case TableNamePolitics:
		return TableLinkPolitics
	case TableNameSociology:
		return TableLinkSociology
	case TableNameTheTempleOfAthena:
		return TableLinkTheTempleOfAthena
	default:
		return ""
	}
}
