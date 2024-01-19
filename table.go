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
	TableNameExistentialism	TableName	= "existentialism"
	TableNamePolitics	TableName	= "politics"
	TableNameSociology	TableName	= "sociology"
	TableNameTheTempleOfAthena	TableName	= "the_temple_of_athena"
	TableNameCommunication	TableName	= "communication"
	TableNameInstrumentality	TableName	= "instrumentality"
	TableNamePyrusMalus	TableName	= "pyrus_malus"
	TableNameJourney	TableName	= "journey"
	TableNameMorality	TableName	= "morality"
	TableNameCulture	TableName	= "culture"
	TableNameIdeology	TableName	= "ideology"
	TableNameNihilism	TableName	= "nihilism"
	TableNameMetaphysics	TableName	= "metaphysics"

	TablePathExistentialism	TablePath	= "/syntelos /science /anthropology /existentialism"
	TablePathPolitics	TablePath	= "/gegonen /politics"
	TablePathSociology	TablePath	= "/syntelos /science /anthropology /sociology"
	TablePathTheTempleOfAthena	TablePath	= "/gegonen /the temple of athena"
	TablePathCommunication	TablePath	= "/syntelos /science /anthropology /communication"
	TablePathInstrumentality	TablePath	= "/syntelos /science /anthropology /instrumentality"
	TablePathPyrusMalus	TablePath	= "/gegonen /through the eyes of pyrus malus"
	TablePathJourney	TablePath	= ""
	TablePathMorality	TablePath	= "/syntelos /science /anthropology /morality"
	TablePathCulture	TablePath	= ""
	TablePathIdeology	TablePath	= "/syntelos /science /anthropology /ideology"
	TablePathNihilism	TablePath	= "/syntelos /science /anthropology /nihilism"
	TablePathMetaphysics	TablePath	= "/syntelos /science /anthropology /metaphysics"

	TableLinkExistentialism	TableLink	= "https://drive.google.com/drive/folders/162g6-KM5dOkWvJX7NgMRoKJuz8rkwk96"
	TableLinkPolitics	TableLink	= "https://drive.google.com/drive/folders/1uhgjgL8HBzRwCgP6pXpCwQqLZ0IC_pgJ"
	TableLinkSociology	TableLink	= "https://drive.google.com/drive/folders/1etCIitYhVQH8_Wf6oHeu_rWUmDfTYZD-"
	TableLinkTheTempleOfAthena	TableLink	= "https://drive.google.com/drive/folders/1nUpMgy9n-wHWN13hlXm-YZH5rojTUi0j"
	TableLinkCommunication	TableLink	= "https://drive.google.com/drive/folders/1tXs2GNe1R9wCsKbj-bcCyeiscrK0F2K0"
	TableLinkInstrumentality	TableLink	= "https://drive.google.com/drive/folders/1YqybUMCurpLc0WdTOuUjew7EujYGFDuG"
	TableLinkPyrusMalus	TableLink	= "https://drive.google.com/drive/folders/1ODny7w7sTRbzQQGYMREZdQaKZ7lHJDuk"
	TableLinkJourney	TableLink	= "/syntelos /science /anthropology /morality"
	TableLinkMorality	TableLink	= "https://drive.google.com/drive/folders/1eLNmCSFIH21y7bWB61QVxw0S3VV9RuZI"
	TableLinkCulture	TableLink	= "/syntelos /science /anthropology /culture"
	TableLinkIdeology	TableLink	= "https://drive.google.com/drive/folders/17CpRzqOoCtS-7WK98N24oPt0Ke_9379E"
	TableLinkNihilism	TableLink	= "https://drive.google.com/drive/folders/1l4CQXKr3DyFJDS3bjfxro7AT5TO_tzZZ"
	TableLinkMetaphysics	TableLink	= "https://drive.google.com/drive/folders/18ea3fgt2FjHWr_HZr84tXCNrmZiQOVHx"
)

func IsTableName(name TableName) bool {
	if 0 != len(name) {
		var path TablePath = name.Path()
		return (0 != len(path))
	} else {
		return false
	}
}

func (this TableName) Path() TablePath {
	switch this {
	case TableNameExistentialism:
		return TablePathExistentialism
	case TableNamePolitics:
		return TablePathPolitics
	case TableNameSociology:
		return TablePathSociology
	case TableNameTheTempleOfAthena:
		return TablePathTheTempleOfAthena
	case TableNameCommunication:
		return TablePathCommunication
	case TableNameInstrumentality:
		return TablePathInstrumentality
	case TableNamePyrusMalus:
		return TablePathPyrusMalus
	case TableNameJourney:
		return TablePathJourney
	case TableNameMorality:
		return TablePathMorality
	case TableNameCulture:
		return TablePathCulture
	case TableNameIdeology:
		return TablePathIdeology
	case TableNameNihilism:
		return TablePathNihilism
	case TableNameMetaphysics:
		return TablePathMetaphysics

	default:
		return ""
	}
}

func (this TableName) Link() TableLink {
	switch this {
	case TableNameExistentialism:
		return TableLinkExistentialism
	case TableNamePolitics:
		return TableLinkPolitics
	case TableNameSociology:
		return TableLinkSociology
	case TableNameTheTempleOfAthena:
		return TableLinkTheTempleOfAthena
	case TableNameCommunication:
		return TableLinkCommunication
	case TableNameInstrumentality:
		return TableLinkInstrumentality
	case TableNamePyrusMalus:
		return TableLinkPyrusMalus
	case TableNameJourney:
		return TableLinkJourney
	case TableNameMorality:
		return TableLinkMorality
	case TableNameCulture:
		return TableLinkCulture
	case TableNameIdeology:
		return TableLinkIdeology
	case TableNameNihilism:
		return TableLinkNihilism
	case TableNameMetaphysics:
		return TableLinkMetaphysics

	default:
		return ""
	}
}
