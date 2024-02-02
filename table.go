/*
 * WWWeb Notes
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package wwweb

type Table string
type TableName Table
type TablePath Table
type TableLink Table

const (

	TableNameBySenseOfReason	TableName	= "by_sense_of_reason"
	TableNameCommunication	TableName	= "communication"
	TableNameCommunism	TableName	= "communism"
	TableNameCulture	TableName	= "culture"
	TableNameEconomics	TableName	= "economics"
	TableNameExistentialism	TableName	= "existentialism"
	TableNameGovernment	TableName	= "government"
	TableNameHumanism	TableName	= "humanism"
	TableNameIdeology	TableName	= "ideology"
	TableNameIndividualism	TableName	= "individualism"
	TableNameInstrumentality	TableName	= "instrumentality"
	TableNameIntellectuality	TableName	= "intellectuality"
	TableNameJourney	TableName	= "journey"
	TableNameLanguage	TableName	= "language"
	TableNameLiberalism	TableName	= "liberalism"
	TableNameLiteracy	TableName	= "literacy"
	TableNameLogic	TableName	= "logic"
	TableNameMaterialism	TableName	= "materialism"
	TableNameMedievalism	TableName	= "medievalism"
	TableNameMetaphysics	TableName	= "metaphysics"
	TableNameMorality	TableName	= "morality"
	TableNameNihilism	TableName	= "nihilism"
	TableNamePlayingWithPowerPolitics	TableName	= "playing_with_power_politics"
	TableNamePolitics	TableName	= "politics"
	TableNamePsychology	TableName	= "psychology"
	TableNamePyrusMalus	TableName	= "pyrus_malus"
	TableNameReflection	TableName	= "reflection"
	TableNameSadism	TableName	= "sadism"
	TableNameSexuality	TableName	= "sexuality"
	TableNameSocialism	TableName	= "socialism"
	TableNameSociology	TableName	= "sociology"
	TableNameSpirituality	TableName	= "spirituality"
	TableNameTheFestivalOfLights	TableName	= "the_festival_of_lights"
	TableNameTheGraceOfAngels	TableName	= "the_grace_of_angels"
	TableNameTheMorrowsPromise	TableName	= "the_morrows_promise"
	TableNameTheMushroomAndTheCrown	TableName	= "the_mushroom_and_the_crown"
	TableNameTheSongOfPsalms	TableName	= "the_song_of_psalms"
	TableNameTheTempleOfAthena	TableName	= "the_temple_of_athena"
	TableNameThroughTheEyesOfPyrusMalus	TableName	= "through_the_eyes_of_pyrus_malus"

	TablePathBySenseOfReason	TablePath	= "/gegonen /through the eyes of pyrus malus /by sense of reason"
	TablePathCommunication	TablePath	= "/syntelos /science /anthropology /communication"
	TablePathCommunism	TablePath	= "/syntelos /science /anthropology /communism"
	TablePathCulture	TablePath	= ""
	TablePathEconomics	TablePath	= "/syntelos /science /anthropology /sociology /economics"
	TablePathExistentialism	TablePath	= "/syntelos /science /anthropology /existentialism"
	TablePathGovernment	TablePath	= "/gegonen /government"
	TablePathHumanism	TablePath	= "/syntelos /science /anthropology /humanism"
	TablePathIdeology	TablePath	= "/syntelos /science /anthropology /ideology"
	TablePathIndividualism	TablePath	= "/syntelos /science /anthropology /individualism"
	TablePathInstrumentality	TablePath	= "/syntelos /science /anthropology /instrumentality"
	TablePathIntellectuality	TablePath	= "/syntelos /science /anthropology /intellectuality"
	TablePathJourney	TablePath	= ""
	TablePathLanguage	TablePath	= "/syntelos /science /anthropology /language"
	TablePathLiberalism	TablePath	= "/syntelos /science /anthropology /liberalism"
	TablePathLiteracy	TablePath	= "/syntelos /science /anthropology /literacy"
	TablePathLogic	TablePath	= "/syntelos /logic"
	TablePathMaterialism	TablePath	= "/syntelos /science /anthropology /materialism"
	TablePathMedievalism	TablePath	= "/gegonen /medievalism"
	TablePathMetaphysics	TablePath	= "/syntelos /science /anthropology /metaphysics"
	TablePathMorality	TablePath	= "/syntelos /science /anthropology /morality"
	TablePathNihilism	TablePath	= "/syntelos /science /anthropology /nihilism"
	TablePathPlayingWithPowerPolitics	TablePath	= "   /gegonen /playing with power politics"
	TablePathPolitics	TablePath	= "/gegonen /politics"
	TablePathPsychology	TablePath	= "/syntelos /science /anthropology /psychology"
	TablePathPyrusMalus	TablePath	= "/gegonen /through the eyes of pyrus malus"
	TablePathReflection	TablePath	= "/gegonen /reflection"
	TablePathSadism	TablePath	= ""
	TablePathSexuality	TablePath	= "/syntelos /science /anthropology /sexuality"
	TablePathSocialism	TablePath	= "/syntelos /science /anthropology /socialism"
	TablePathSociology	TablePath	= "/syntelos /science /anthropology /sociology"
	TablePathSpirituality	TablePath	= "/syntelos /science /anthropology /spirituality"
	TablePathTheFestivalOfLights	TablePath	= "/gegonen /through the eyes of pyrus malus /the festival of lights"
	TablePathTheGraceOfAngels	TablePath	= "/gegonen /through the eyes of pyrus malus /the grace of angels"
	TablePathTheMorrowsPromise	TablePath	= "/gegonen /through the eyes of pyrus malus /the morrows promise"
	TablePathTheMushroomAndTheCrown	TablePath	= "/gegonen /through the eyes of pyrus malus /the mushroom and the crown"
	TablePathTheSongOfPsalms	TablePath	= "/gegonen /through the eyes of pyrus malus /the song of psalsm"
	TablePathTheTempleOfAthena	TablePath	= "/gegonen /the temple of athena"
	TablePathThroughTheEyesOfPyrusMalus	TablePath	= "/gegonen /through the eyes of pyrus malus"

	TableLinkBySenseOfReason	TableLink	= "https://drive.google.com/drive/folders/1u1yE4jdskWzKoEuJXFPUoLIXnwj6KISv"
	TableLinkCommunication	TableLink	= "https://drive.google.com/drive/folders/1tXs2GNe1R9wCsKbj-bcCyeiscrK0F2K0"
	TableLinkCommunism	TableLink	= "https://drive.google.com/drive/folders/15Eq2asOsXyz9xB-A-NLFU1IUg0lKDQWJ"
	TableLinkCulture	TableLink	= "/syntelos /science /anthropology /culture"
	TableLinkEconomics	TableLink	= "https://drive.google.com/drive/folders/11KR2sEIWVrBCcWzNcTsI1PBRLMpykzrB"
	TableLinkExistentialism	TableLink	= "https://drive.google.com/drive/folders/162g6-KM5dOkWvJX7NgMRoKJuz8rkwk96"
	TableLinkGovernment	TableLink	= "https://drive.google.com/drive/folders/1uhgjgL8HBzRwCgP6pXpCwQqLZ0IC_pgJ"
	TableLinkHumanism	TableLink	= "https://drive.google.com/drive/folders/1QHhQnpt88m0u5dN2fgHto8Dam5Mh6c_9"
	TableLinkIdeology	TableLink	= "https://drive.google.com/drive/folders/17CpRzqOoCtS-7WK98N24oPt0Ke_9379E"
	TableLinkIndividualism	TableLink	= "https://drive.google.com/drive/folders/1vahEEfMA42RpSTXNEZgEMeuO7SDOqXg0"
	TableLinkInstrumentality	TableLink	= "https://drive.google.com/drive/folders/1YqybUMCurpLc0WdTOuUjew7EujYGFDuG"
	TableLinkIntellectuality	TableLink	= "https://drive.google.com/drive/folders/1_zMHs6uHkF4UzlIywjqwfv9FmZyGO5rT"
	TableLinkJourney	TableLink	= "/syntelos /science /anthropology /morality"
	TableLinkLanguage	TableLink	= "https://drive.google.com/drive/folders/1B5PyRFZDxYQnqBP39kO1C-BlV7O2PD46"
	TableLinkLiberalism	TableLink	= "https://drive.google.com/drive/folders/1OczHAQbfNIOwBzhqwCFoCPPQM_eDsQ1J"
	TableLinkLiteracy	TableLink	= "https://drive.google.com/drive/folders/1KeSxi5xwUFB-D51PolHoGXEfZ2fLvwMo"
	TableLinkLogic	TableLink	= "https://drive.google.com/drive/folders/1YIQWS_9QQwhX0TGkliZ5bXVN74nqAPI8"
	TableLinkMaterialism	TableLink	= "https://drive.google.com/drive/folders/1KbBCIIRH5RtqyPda8NW14dcRQh2fwsvL"
	TableLinkMedievalism	TableLink	= "https://drive.google.com/drive/folders/1I8AYk_tCErT61UG6DCArM3Pbt1gyzw4s"
	TableLinkMetaphysics	TableLink	= "https://drive.google.com/drive/folders/18ea3fgt2FjHWr_HZr84tXCNrmZiQOVHx"
	TableLinkMorality	TableLink	= "https://drive.google.com/drive/folders/1eLNmCSFIH21y7bWB61QVxw0S3VV9RuZI"
	TableLinkNihilism	TableLink	= "https://drive.google.com/drive/folders/1l4CQXKr3DyFJDS3bjfxro7AT5TO_tzZZ"
	TableLinkPlayingWithPowerPolitics	TableLink	= "https://drive.google.com/drive/folders/1EcXQM3IhogaEFpoLBpUoQJwnD9fajhGi"
	TableLinkPolitics	TableLink	= "https://drive.google.com/drive/folders/1uhgjgL8HBzRwCgP6pXpCwQqLZ0IC_pgJ"
	TableLinkPsychology	TableLink	= "https://drive.google.com/drive/folders/1adSWRCYYCf_rdhcxUk2hCe-k3Tn6tzHQ"
	TableLinkPyrusMalus	TableLink	= "https://drive.google.com/drive/folders/1ODny7w7sTRbzQQGYMREZdQaKZ7lHJDuk"
	TableLinkReflection	TableLink	= "https://drive.google.com/drive/folders/10-TWr2dAPdA60uh3EyLpCIs0Kcy3j_vZ"
	TableLinkSadism	TableLink	= "/syntelos /science /anthropology /nihilism /sadism"
	TableLinkSexuality	TableLink	= "https://drive.google.com/drive/folders/10EU25uM8ueRaLXI_WT0PtAE4nRR6tRTb"
	TableLinkSocialism	TableLink	= "https://drive.google.com/drive/folders/1v6n1KFL4pGDRZo-LmfkHBTs0k-QUTMbj"
	TableLinkSociology	TableLink	= "https://drive.google.com/drive/folders/1etCIitYhVQH8_Wf6oHeu_rWUmDfTYZD-"
	TableLinkSpirituality	TableLink	= "https://drive.google.com/drive/folders/1Ztj8tmFI6qvkVkjx5Y62VqopK_6Sq86x"
	TableLinkTheFestivalOfLights	TableLink	= "https://drive.google.com/drive/folders/1l9Ceck-1kcwJ3G0_sSsB2xgWa8SxBQxn"
	TableLinkTheGraceOfAngels	TableLink	= "https://drive.google.com/drive/folders/1d3jFpOaoUHQHJP1oQGiZXNb3tLrRwqas"
	TableLinkTheMorrowsPromise	TableLink	= "https://drive.google.com/drive/folders/13MZnnKVVUuR8zZLPFZGq4oOM_pFGmIGw"
	TableLinkTheMushroomAndTheCrown	TableLink	= "https://drive.google.com/drive/folders/1tyaz1GXRo0HfJMZR2rAzw6tZLpAeIjX7"
	TableLinkTheSongOfPsalms	TableLink	= "https://drive.google.com/drive/folders/1nOX1MerIv2rULUJq-BKpV4ZKF3jexyMD"
	TableLinkTheTempleOfAthena	TableLink	= "https://drive.google.com/drive/folders/1nUpMgy9n-wHWN13hlXm-YZH5rojTUi0j"
	TableLinkThroughTheEyesOfPyrusMalus	TableLink	= "https://drive.google.com/drive/folders/1ODny7w7sTRbzQQGYMREZdQaKZ7lHJDuk"
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

	case TableNameBySenseOfReason:
		return TablePathBySenseOfReason
	case TableNameCommunication:
		return TablePathCommunication
	case TableNameCommunism:
		return TablePathCommunism
	case TableNameCulture:
		return TablePathCulture
	case TableNameEconomics:
		return TablePathEconomics
	case TableNameExistentialism:
		return TablePathExistentialism
	case TableNameGovernment:
		return TablePathGovernment
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
	case TableNameLanguage:
		return TablePathLanguage
	case TableNameLiberalism:
		return TablePathLiberalism
	case TableNameLiteracy:
		return TablePathLiteracy
	case TableNameLogic:
		return TablePathLogic
	case TableNameMaterialism:
		return TablePathMaterialism
	case TableNameMedievalism:
		return TablePathMedievalism
	case TableNameMetaphysics:
		return TablePathMetaphysics
	case TableNameMorality:
		return TablePathMorality
	case TableNameNihilism:
		return TablePathNihilism
	case TableNamePlayingWithPowerPolitics:
		return TablePathPlayingWithPowerPolitics
	case TableNamePolitics:
		return TablePathPolitics
	case TableNamePsychology:
		return TablePathPsychology
	case TableNamePyrusMalus:
		return TablePathPyrusMalus
	case TableNameReflection:
		return TablePathReflection
	case TableNameSadism:
		return TablePathSadism
	case TableNameSexuality:
		return TablePathSexuality
	case TableNameSocialism:
		return TablePathSocialism
	case TableNameSociology:
		return TablePathSociology
	case TableNameSpirituality:
		return TablePathSpirituality
	case TableNameTheFestivalOfLights:
		return TablePathTheFestivalOfLights
	case TableNameTheGraceOfAngels:
		return TablePathTheGraceOfAngels
	case TableNameTheMorrowsPromise:
		return TablePathTheMorrowsPromise
	case TableNameTheMushroomAndTheCrown:
		return TablePathTheMushroomAndTheCrown
	case TableNameTheSongOfPsalms:
		return TablePathTheSongOfPsalms
	case TableNameTheTempleOfAthena:
		return TablePathTheTempleOfAthena
	case TableNameThroughTheEyesOfPyrusMalus:
		return TablePathThroughTheEyesOfPyrusMalus

	default:
		return ""
	}
}

func (this TableName) Link() TableLink {
	switch this {

	case TableNameBySenseOfReason:
		return TableLinkBySenseOfReason
	case TableNameCommunication:
		return TableLinkCommunication
	case TableNameCommunism:
		return TableLinkCommunism
	case TableNameCulture:
		return TableLinkCulture
	case TableNameEconomics:
		return TableLinkEconomics
	case TableNameExistentialism:
		return TableLinkExistentialism
	case TableNameGovernment:
		return TableLinkGovernment
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
	case TableNameLanguage:
		return TableLinkLanguage
	case TableNameLiberalism:
		return TableLinkLiberalism
	case TableNameLiteracy:
		return TableLinkLiteracy
	case TableNameLogic:
		return TableLinkLogic
	case TableNameMaterialism:
		return TableLinkMaterialism
	case TableNameMedievalism:
		return TableLinkMedievalism
	case TableNameMetaphysics:
		return TableLinkMetaphysics
	case TableNameMorality:
		return TableLinkMorality
	case TableNameNihilism:
		return TableLinkNihilism
	case TableNamePlayingWithPowerPolitics:
		return TableLinkPlayingWithPowerPolitics
	case TableNamePolitics:
		return TableLinkPolitics
	case TableNamePsychology:
		return TableLinkPsychology
	case TableNamePyrusMalus:
		return TableLinkPyrusMalus
	case TableNameReflection:
		return TableLinkReflection
	case TableNameSadism:
		return TableLinkSadism
	case TableNameSexuality:
		return TableLinkSexuality
	case TableNameSocialism:
		return TableLinkSocialism
	case TableNameSociology:
		return TableLinkSociology
	case TableNameSpirituality:
		return TableLinkSpirituality
	case TableNameTheFestivalOfLights:
		return TableLinkTheFestivalOfLights
	case TableNameTheGraceOfAngels:
		return TableLinkTheGraceOfAngels
	case TableNameTheMorrowsPromise:
		return TableLinkTheMorrowsPromise
	case TableNameTheMushroomAndTheCrown:
		return TableLinkTheMushroomAndTheCrown
	case TableNameTheSongOfPsalms:
		return TableLinkTheSongOfPsalms
	case TableNameTheTempleOfAthena:
		return TableLinkTheTempleOfAthena
	case TableNameThroughTheEyesOfPyrusMalus:
		return TableLinkThroughTheEyesOfPyrusMalus

	default:
		return ""
	}
}

