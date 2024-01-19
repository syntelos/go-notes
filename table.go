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
	TableNameCommunication	TableName	= "communication"
	TableNameCommunism	TableName	= "communism"
	TableNameCulture	TableName	= "culture"
	TableNameExistentialism	TableName	= "existentialism"
	TableNameHumanism	TableName	= "humanism"
	TableNameIdeology	TableName	= "ideology"
	TableNameIndividualism	TableName	= "individualism"
	TableNameInstrumentality	TableName	= "instrumentality"
	TableNameIntellectuality	TableName	= "intellectuality"
	TableNameJourney	TableName	= "journey"
	TableNameLiberalism	TableName	= "liberalism"
	TableNameLiteracy	TableName	= "literacy"
	TableNameMaterialism	TableName	= "materialism"
	TableNameMetaphysics	TableName	= "metaphysics"
	TableNameMorality	TableName	= "morality"
	TableNameNihilism	TableName	= "nihilism"
	TableNamePolitics	TableName	= "politics"
	TableNamePsychology	TableName	= "psychology"
	TableNamePyrusMalus	TableName	= "pyrus_malus"
	TableNameSexuality	TableName	= "sexuality"
	TableNameSocialism	TableName	= "socialism"
	TableNameSociology	TableName	= "sociology"
	TableNameSpirituality	TableName	= "spirituality"
	TableNameTheTempleOfAthena	TableName	= "the_temple_of_athena"

	TablePathCommunication	TablePath	= "/syntelos /science /anthropology /communication"
	TablePathCommunism	TablePath	= "/syntelos /science /anthropology /communism"
	TablePathCulture	TablePath	= ""
	TablePathExistentialism	TablePath	= "/syntelos /science /anthropology /existentialism"
	TablePathHumanism	TablePath	= "/syntelos /science /anthropology /humanism"
	TablePathIdeology	TablePath	= "/syntelos /science /anthropology /ideology"
	TablePathIndividualism	TablePath	= "/syntelos /science /anthropology /individualism"
	TablePathInstrumentality	TablePath	= "/syntelos /science /anthropology /instrumentality"
	TablePathIntellectuality	TablePath	= "/syntelos /science /anthropology /intellectuality"
	TablePathJourney	TablePath	= ""
	TablePathLiberalism	TablePath	= "/syntelos /science /anthropology /liberalism"
	TablePathLiteracy	TablePath	= "/syntelos /science /anthropology /literacy"
	TablePathMaterialism	TablePath	= "/syntelos /science /anthropology /materialism"
	TablePathMetaphysics	TablePath	= "/syntelos /science /anthropology /metaphysics"
	TablePathMorality	TablePath	= "/syntelos /science /anthropology /morality"
	TablePathNihilism	TablePath	= "/syntelos /science /anthropology /nihilism"
	TablePathPolitics	TablePath	= "/gegonen /politics"
	TablePathPsychology	TablePath	= "/syntelos /science /anthropology /psychology"
	TablePathPyrusMalus	TablePath	= "/gegonen /through the eyes of pyrus malus"
	TablePathSexuality	TablePath	= "/syntelos /science /anthropology /sexuality"
	TablePathSocialism	TablePath	= "/syntelos /science /anthropology /socialism"
	TablePathSociology	TablePath	= "/syntelos /science /anthropology /sociology"
	TablePathSpirituality	TablePath	= "/syntelos /science /anthropology /spirituality"
	TablePathTheTempleOfAthena	TablePath	= "/gegonen /the temple of athena"

	TableLinkCommunication	TableLink	= "https://drive.google.com/drive/folders/1tXs2GNe1R9wCsKbj-bcCyeiscrK0F2K0"
	TableLinkCommunism	TableLink	= "https://drive.google.com/drive/folders/15Eq2asOsXyz9xB-A-NLFU1IUg0lKDQWJ"
	TableLinkCulture	TableLink	= "/syntelos /science /anthropology /culture"
	TableLinkExistentialism	TableLink	= "https://drive.google.com/drive/folders/162g6-KM5dOkWvJX7NgMRoKJuz8rkwk96"
	TableLinkHumanism	TableLink	= "https://drive.google.com/drive/folders/1QHhQnpt88m0u5dN2fgHto8Dam5Mh6c_9"
	TableLinkIdeology	TableLink	= "https://drive.google.com/drive/folders/17CpRzqOoCtS-7WK98N24oPt0Ke_9379E"
	TableLinkIndividualism	TableLink	= "https://drive.google.com/drive/folders/1vahEEfMA42RpSTXNEZgEMeuO7SDOqXg0"
	TableLinkInstrumentality	TableLink	= "https://drive.google.com/drive/folders/1YqybUMCurpLc0WdTOuUjew7EujYGFDuG"
	TableLinkIntellectuality	TableLink	= "https://drive.google.com/drive/folders/1_zMHs6uHkF4UzlIywjqwfv9FmZyGO5rT"
	TableLinkJourney	TableLink	= "/syntelos /science /anthropology /morality"
	TableLinkLiberalism	TableLink	= "https://drive.google.com/drive/folders/1OczHAQbfNIOwBzhqwCFoCPPQM_eDsQ1J"
	TableLinkLiteracy	TableLink	= "https://drive.google.com/drive/folders/1KeSxi5xwUFB-D51PolHoGXEfZ2fLvwMo"
	TableLinkMaterialism	TableLink	= "https://drive.google.com/drive/folders/1KbBCIIRH5RtqyPda8NW14dcRQh2fwsvL"
	TableLinkMetaphysics	TableLink	= "https://drive.google.com/drive/folders/18ea3fgt2FjHWr_HZr84tXCNrmZiQOVHx"
	TableLinkMorality	TableLink	= "https://drive.google.com/drive/folders/1eLNmCSFIH21y7bWB61QVxw0S3VV9RuZI"
	TableLinkNihilism	TableLink	= "https://drive.google.com/drive/folders/1l4CQXKr3DyFJDS3bjfxro7AT5TO_tzZZ"
	TableLinkPolitics	TableLink	= "https://drive.google.com/drive/folders/1uhgjgL8HBzRwCgP6pXpCwQqLZ0IC_pgJ"
	TableLinkPsychology	TableLink	= "https://drive.google.com/drive/folders/1adSWRCYYCf_rdhcxUk2hCe-k3Tn6tzHQ"
	TableLinkPyrusMalus	TableLink	= "https://drive.google.com/drive/folders/1ODny7w7sTRbzQQGYMREZdQaKZ7lHJDuk"
	TableLinkSexuality	TableLink	= "https://drive.google.com/drive/folders/10EU25uM8ueRaLXI_WT0PtAE4nRR6tRTb"
	TableLinkSocialism	TableLink	= "https://drive.google.com/drive/folders/1v6n1KFL4pGDRZo-LmfkHBTs0k-QUTMbj"
	TableLinkSociology	TableLink	= "https://drive.google.com/drive/folders/1etCIitYhVQH8_Wf6oHeu_rWUmDfTYZD-"
	TableLinkSpirituality	TableLink	= "https://drive.google.com/drive/folders/1Ztj8tmFI6qvkVkjx5Y62VqopK_6Sq86x"
	TableLinkTheTempleOfAthena	TableLink	= "https://drive.google.com/drive/folders/1nUpMgy9n-wHWN13hlXm-YZH5rojTUi0j"
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
	case TableNameCommunication:
		return TablePathCommunication
	case TableNameCommunism:
		return TablePathCommunism
	case TableNameCulture:
		return TablePathCulture
	case TableNameExistentialism:
		return TablePathExistentialism
	case TableNameHumanism:
		return TablePathHumanism
	case TableNameIdeology:
		return TablePathIdeology
	case TableNameIndividualism:
		return TablePathIndividualism
	case TableNameInstrumentality:
		return TablePathInstrumentality
	case TableNameIntellectuality:
		return TablePathIntellectuality
	case TableNameJourney:
		return TablePathJourney
	case TableNameLiberalism:
		return TablePathLiberalism
	case TableNameLiteracy:
		return TablePathLiteracy
	case TableNameMaterialism:
		return TablePathMaterialism
	case TableNameMetaphysics:
		return TablePathMetaphysics
	case TableNameMorality:
		return TablePathMorality
	case TableNameNihilism:
		return TablePathNihilism
	case TableNamePolitics:
		return TablePathPolitics
	case TableNamePsychology:
		return TablePathPsychology
	case TableNamePyrusMalus:
		return TablePathPyrusMalus
	case TableNameSexuality:
		return TablePathSexuality
	case TableNameSocialism:
		return TablePathSocialism
	case TableNameSociology:
		return TablePathSociology
	case TableNameSpirituality:
		return TablePathSpirituality
	case TableNameTheTempleOfAthena:
		return TablePathTheTempleOfAthena

	default:
		return ""
	}
}

func (this TableName) Link() TableLink {
	switch this {
	case TableNameCommunication:
		return TableLinkCommunication
	case TableNameCommunism:
		return TableLinkCommunism
	case TableNameCulture:
		return TableLinkCulture
	case TableNameExistentialism:
		return TableLinkExistentialism
	case TableNameHumanism:
		return TableLinkHumanism
	case TableNameIdeology:
		return TableLinkIdeology
	case TableNameIndividualism:
		return TableLinkIndividualism
	case TableNameInstrumentality:
		return TableLinkInstrumentality
	case TableNameIntellectuality:
		return TableLinkIntellectuality
	case TableNameJourney:
		return TableLinkJourney
	case TableNameLiberalism:
		return TableLinkLiberalism
	case TableNameLiteracy:
		return TableLinkLiteracy
	case TableNameMaterialism:
		return TableLinkMaterialism
	case TableNameMetaphysics:
		return TableLinkMetaphysics
	case TableNameMorality:
		return TableLinkMorality
	case TableNameNihilism:
		return TableLinkNihilism
	case TableNamePolitics:
		return TableLinkPolitics
	case TableNamePsychology:
		return TableLinkPsychology
	case TableNamePyrusMalus:
		return TableLinkPyrusMalus
	case TableNameSexuality:
		return TableLinkSexuality
	case TableNameSocialism:
		return TableLinkSocialism
	case TableNameSociology:
		return TableLinkSociology
	case TableNameSpirituality:
		return TableLinkSpirituality
	case TableNameTheTempleOfAthena:
		return TableLinkTheTempleOfAthena

	default:
		return ""
	}
}
