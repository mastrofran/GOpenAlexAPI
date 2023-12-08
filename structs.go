package OpenAlexAPI

type Work struct {
	ID     string
	Select []string
}

type Client struct {
	MailTo  *string
	BaseURL string
}

type Response struct {
	ID              *string `json:"id"`
	Doi             *string `json:"doi"`
	Title           *string `json:"title"`
	DisplayName     *string `json:"display_name,omitempty"`
	PublicationYear *int    `json:"publication_year,omitempty"`
	PublicationDate *string `json:"publication_date,omitempty"`
	Ids             struct {
		Openalex string `json:"openalex"`
		Doi      string `json:"doi"`
		Mag      string `json:"mag"`
	} `json:"ids"`
	Language        string `json:"language"`
	PrimaryLocation struct {
		IsOa           bool   `json:"is_oa"`
		LandingPageURL string `json:"landing_page_url"`
		PdfURL         any    `json:"pdf_url"`
		Source         struct {
			ID                           string   `json:"id"`
			DisplayName                  string   `json:"display_name"`
			IssnL                        string   `json:"issn_l"`
			Issn                         []string `json:"issn"`
			IsOa                         bool     `json:"is_oa"`
			IsInDoaj                     bool     `json:"is_in_doaj"`
			HostOrganization             string   `json:"host_organization"`
			HostOrganizationName         string   `json:"host_organization_name"`
			HostOrganizationLineage      []string `json:"host_organization_lineage"`
			HostOrganizationLineageNames []string `json:"host_organization_lineage_names"`
			Type                         string   `json:"type"`
		} `json:"source"`
		License     string `json:"license"`
		Version     string `json:"version"`
		IsAccepted  bool   `json:"is_accepted"`
		IsPublished bool   `json:"is_published"`
	} `json:"primary_location"`
	Type         string `json:"type"`
	TypeCrossref string `json:"type_crossref"`
	OpenAccess   struct {
		IsOa                     bool   `json:"is_oa"`
		OaStatus                 string `json:"oa_status"`
		OaURL                    string `json:"oa_url"`
		AnyRepositoryHasFulltext bool   `json:"any_repository_has_fulltext"`
	} `json:"open_access"`
	Authorships []struct {
		AuthorPosition string `json:"author_position"`
		Author         struct {
			ID          string `json:"id"`
			DisplayName string `json:"display_name"`
			Orcid       any    `json:"orcid"`
		} `json:"author"`
		Institutions []struct {
			ID          string   `json:"id"`
			DisplayName string   `json:"display_name"`
			Ror         string   `json:"ror"`
			CountryCode string   `json:"country_code"`
			Type        string   `json:"type"`
			Lineage     []string `json:"lineage"`
		} `json:"institutions"`
		Countries             []string `json:"countries"`
		IsCorresponding       bool     `json:"is_corresponding"`
		RawAuthorName         string   `json:"raw_author_name"`
		RawAffiliationString  string   `json:"raw_affiliation_string"`
		RawAffiliationStrings []string `json:"raw_affiliation_strings"`
	} `json:"authorships"`
	CountriesDistinctCount      int   `json:"countries_distinct_count"`
	InstitutionsDistinctCount   int   `json:"institutions_distinct_count"`
	CorrespondingAuthorIds      []any `json:"corresponding_author_ids"`
	CorrespondingInstitutionIds []any `json:"corresponding_institution_ids"`
	ApcList                     struct {
		Value      int    `json:"value"`
		Currency   string `json:"currency"`
		ValueUsd   int    `json:"value_usd"`
		Provenance string `json:"provenance"`
	} `json:"apc_list"`
	ApcPaid struct {
		Value      int    `json:"value"`
		Currency   string `json:"currency"`
		ValueUsd   int    `json:"value_usd"`
		Provenance string `json:"provenance"`
	} `json:"apc_paid"`
	HasFulltext           bool   `json:"has_fulltext"`
	FulltextOrigin        string `json:"fulltext_origin"`
	CitedByCount          int    `json:"cited_by_count"`
	CitedByPercentileYear struct {
		Min float64 `json:"min"`
		Max float64 `json:"max"`
	} `json:"cited_by_percentile_year"`
	Biblio struct {
		Volume    string `json:"volume"`
		Issue     string `json:"issue"`
		FirstPage string `json:"first_page"`
		LastPage  string `json:"last_page"`
	} `json:"biblio"`
	IsRetracted bool `json:"is_retracted"`
	IsParatext  bool `json:"is_paratext"`
	Keywords    []struct {
		Keyword string  `json:"keyword"`
		Score   float64 `json:"score"`
	} `json:"keywords"`
	Concepts []struct {
		ID          string  `json:"id"`
		Wikidata    string  `json:"wikidata"`
		DisplayName string  `json:"display_name"`
		Level       int     `json:"level"`
		Score       float64 `json:"score"`
	} `json:"concepts"`
	Mesh           []any `json:"mesh"`
	LocationsCount int   `json:"locations_count"`
	Locations      []struct {
		IsOa           bool   `json:"is_oa"`
		LandingPageURL string `json:"landing_page_url"`
		PdfURL         any    `json:"pdf_url"`
		Source         struct {
			ID                           string   `json:"id"`
			DisplayName                  string   `json:"display_name"`
			IssnL                        string   `json:"issn_l"`
			Issn                         []string `json:"issn"`
			IsOa                         bool     `json:"is_oa"`
			IsInDoaj                     bool     `json:"is_in_doaj"`
			HostOrganization             string   `json:"host_organization"`
			HostOrganizationName         string   `json:"host_organization_name"`
			HostOrganizationLineage      []string `json:"host_organization_lineage"`
			HostOrganizationLineageNames []string `json:"host_organization_lineage_names"`
			Type                         string   `json:"type"`
		} `json:"source"`
		License     string `json:"license"`
		Version     string `json:"version"`
		IsAccepted  bool   `json:"is_accepted"`
		IsPublished bool   `json:"is_published"`
	} `json:"locations"`
	BestOaLocation struct {
		IsOa           bool   `json:"is_oa"`
		LandingPageURL string `json:"landing_page_url"`
		PdfURL         any    `json:"pdf_url"`
		Source         struct {
			ID                           string   `json:"id"`
			DisplayName                  string   `json:"display_name"`
			IssnL                        string   `json:"issn_l"`
			Issn                         []string `json:"issn"`
			IsOa                         bool     `json:"is_oa"`
			IsInDoaj                     bool     `json:"is_in_doaj"`
			HostOrganization             string   `json:"host_organization"`
			HostOrganizationName         string   `json:"host_organization_name"`
			HostOrganizationLineage      []string `json:"host_organization_lineage"`
			HostOrganizationLineageNames []string `json:"host_organization_lineage_names"`
			Type                         string   `json:"type"`
		} `json:"source"`
		License     string `json:"license"`
		Version     string `json:"version"`
		IsAccepted  bool   `json:"is_accepted"`
		IsPublished bool   `json:"is_published"`
	} `json:"best_oa_location"`
	SustainableDevelopmentGoals []struct {
		ID          string  `json:"id"`
		DisplayName string  `json:"display_name"`
		Score       float64 `json:"score"`
	} `json:"sustainable_development_goals"`
	Grants                []any    `json:"grants"`
	ReferencedWorksCount  int      `json:"referenced_works_count"`
	ReferencedWorks       []string `json:"referenced_works"`
	RelatedWorks          []string `json:"related_works"`
	NgramsURL             string   `json:"ngrams_url"`
	AbstractInvertedIndex struct {
		Num100        []int `json:"100"`
		Summary       []int `json:"Summary"`
		OneNew        []int `json:"1.New"`
		Techniques    []int `json:"techniques"`
		For           []int `json:"for"`
		Studying      []int `json:"studying"`
		The           []int `json:"the"`
		Microbiology  []int `json:"microbiology"`
		Of            []int `json:"of"`
		Rumen         []int `json:"rumen"`
		Are           []int `json:"are"`
		Described     []int `json:"described."`
		These         []int `json:"These"`
		Include       []int `json:"include"`
		A             []int `json:"a"`
		Direct        []int `json:"direct"`
		Slide         []int `json:"slide"`
		Count         []int `json:"count"`
		Technique     []int `json:"technique,"`
		Cultural      []int `json:"cultural"`
		Procedure     []int `json:"procedure"`
		Isolating     []int `json:"isolating"`
		Predominating []int `json:"predominating"`
		Flora         []int `json:"flora,"`
		Method        []int `json:"method"`
		Obtaining     []int `json:"obtaining"`
		Sample        []int `json:"sample"`
		Under         []int `json:"under"`
		Anaerobic     []int `json:"anaerobic"`
		Conditions    []int `json:"conditions,"`
		And           []int `json:"and"`
		Methods       []int `json:"methods"`
		Determination []int `json:"determination"`
		Urea          []int `json:"urea"`
		Utilization   []int `json:"utilization"`
		Cellulose     []int `json:"cellulose"`
		Digestion     []int `json:"digestion"`
		By            []int `json:"by"`
		Bacteria      []int `json:"bacteria."`
		TwoThe        []int `json:"2.The"`
		Technique0    []int `json:"technique"`
		Made          []int `json:"made"`
		Use           []int `json:"use"`
		Nigrosine     []int `json:"nigrosine,"`
		Negative      []int `json:"negative"`
		Stain         []int `json:"stain,"`
		Which         []int `json:"which"`
		Reduced       []int `json:"reduced"`
		To            []int `json:"to"`
		Minimum       []int `json:"minimum"`
		Inaccuracies  []int `json:"inaccuracies"`
		Caused        []int `json:"caused"`
		Artifacts     []int `json:"artifacts."`
		ThreeWhen     []int `json:"3.When"`
		Good          []int `json:"good"`
		Conditions0   []int `json:"conditions"`
		Prevailed     []int `json:"prevailed"`
		Fistulated    []int `json:"fistulated"`
		Animal        []int `json:"animal"`
		Was           []int `json:"was"`
		Used          []int `json:"used,"`
		Counts        []int `json:"counts"`
		Agreed        []int `json:"agreed"`
		Closely       []int `json:"closely."`
		Material      []int `json:"Material"`
		Obtained      []int `json:"obtained"`
		Stomach       []int `json:"stomach"`
		Tube          []int `json:"tube"`
		Gave          []int `json:"gave"`
		Consistently  []int `json:"consistently"`
		Lower         []int `json:"lower"`
		Bacterial     []int `json:"bacterial"`
		Than          []int `json:"than"`
		Samples       []int `json:"samples"`
		Taken         []int `json:"taken"`
		Directly      []int `json:"directly"`
		From          []int `json:"from"`
		Fistula       []int `json:"fistula."`
		To0           []int `json:"To"`
		Isolate       []int `json:"isolate"`
		Most          []int `json:"most"`
		Important     []int `json:"important"`
		Feature       []int `json:"feature"`
		Maintenance   []int `json:"maintenance"`
		Time          []int `json:"time"`
		Sampling      []int `json:"sampling"`
		Until         []int `json:"until"`
		Growth        []int `json:"growth"`
		Obtained0     []int `json:"obtained."`
		A0            []int `json:"A"`
		Rich          []int `json:"rich"`
		Organic       []int `json:"organic"`
		Medium        []int `json:"medium"`
		With          []int `json:"with"`
		PH            []int `json:"pH"`
		Six0          []int `json:"6.0"`
		Used0         []int `json:"used"`
		Cultures      []int `json:"cultures"`
		Were          []int `json:"were"`
		Incubated     []int `json:"incubated"`
		At            []int `json:"at"`
		Three8        []int `json:"38°"`
		C             []int `json:"C."`
		FourIt        []int `json:"4.It"`
		Shown         []int `json:"shown"`
		These0        []int `json:"these"`
		That          []int `json:"that"`
		Bacteria0     []int `json:"bacteria"`
		Present       []int `json:"present"`
		In            []int `json:"in"`
		Contents      []int `json:"contents"`
		Both          []int `json:"both"`
		Sheep         []int `json:"sheep"`
		Cattle        []int `json:"cattle"`
		Numbers       []int `json:"numbers"`
		About         []int `json:"about"`
		Billion       []int `json:"billion"`
		Per           []int `json:"per"`
		Gram          []int `json:"gram"`
		Fresh         []int `json:"fresh"`
		Material0     []int `json:"material."`
		FiveBacteria  []int `json:"5.Bacteria"`
		Isolated      []int `json:"isolated"`
		Highest       []int `json:"highest"`
		Dilutions     []int `json:"dilutions"`
		Tested        []int `json:"tested"`
		Their         []int `json:"their"`
		Ability       []int `json:"ability"`
		Utilize       []int `json:"utilize"`
		Break         []int `json:"break"`
		Down          []int `json:"down"`
		Cellulose0    []int `json:"cellulose,"`
		Many          []int `json:"many"`
		Proved        []int `json:"proved"`
		Capable       []int `json:"capable"`
		Performing    []int `json:"performing"`
		Functions     []int `json:"functions"`
		Vitro         []int `json:"vitro."`
		SixA          []int `json:"6.A"`
		Few           []int `json:"few"`
		Studied       []int `json:"studied"`
		Found         []int `json:"found"`
		Be            []int `json:"be"`
		Mostly        []int `json:"mostly"`
		GramPositive  []int `json:"Gram-positive,"`
		NonMotile     []int `json:"non-motile"`
		Rods          []int `json:"rods."`
		Small         []int `json:"small"`
		Number        []int `json:"number"`
		Cocci         []int `json:"cocci"`
		Cultured      []int `json:"cultured."`
	} `json:"abstract_inverted_index"`
	CitedByAPIURL string `json:"cited_by_api_url"`
	CountsByYear  []struct {
		Year         int `json:"year"`
		CitedByCount int `json:"cited_by_count"`
	} `json:"counts_by_year"`
	UpdatedDate string `json:"updated_date"`
	CreatedDate string `json:"created_date"`
}
