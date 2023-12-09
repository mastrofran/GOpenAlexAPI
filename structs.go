package GOpenAlexAPI

type WorkQuery struct {
	ID      string
	Select  []string
	PerPage int
	PageNum int
	Sample  int
	Seed    int
	Filter  string
	Search  string
	Sort    string
}

type Client struct {
	MailTo  *string
	BaseURL string
}

type Work struct {
	AbstractInvertedIndex map[string][]int `json:"abstract_inverted_index,omitempty"`
	Authorships           []struct {
		AuthorPosition string   `json:"author_position,omitempty"`
		Countries      []string `json:"countries,omitempty"`
		Author         struct {
			ID          string `json:"id,omitempty"`
			DisplayName string `json:"display_name,omitempty"`
			ORCID       string `json:"ORCID,omitempty"`
		}
		Institutions []struct {
			ID          string   `json:"id,omitempty"`
			DisplayName string   `json:"display_name,omitempty"`
			Ror         string   `json:"ror,omitempty"`
			CountryCode string   `json:"country_code,omitempty"`
			Type        string   `json:"type,omitempty"`
			Lineage     []string `json:"lineage,omitempty"`
		} `json:"institutions,omitempty"`
		IsCorresponding      bool   `json:"is_corresponding,omitempty"`
		RawAffiliationString string `json:"raw_affiliation_string,omitempty"`
		RawAuthorName        string `json:"raw_author_name,omitempty"`
	} `json:"authorships,omitempty"`
	APCList struct {
		Value      int    `json:"value,omitempty"`
		Currency   string `json:"currency,omitempty"`
		Provenance string `json:"provenance,omitempty"`
		ValueUSD   int    `json:"value_usd,omitempty"`
	} `json:"apc_list,omitempty"`
	APCPaid struct {
		Value      int    `json:"value,omitempty"`
		Currency   string `json:"currency,omitempty"`
		Provenance string `json:"provenance,omitempty"`
		ValueUSD   int    `json:"value_usd,omitempty"`
	} `json:"apc_paid,omitempty"`
	BestOALocation struct {
		IsAccepted     bool   `json:"is_accepted,omitempty"`
		IsOA           bool   `json:"is_oa,omitempty"`
		IsPublished    bool   `json:"is_published,omitempty"`
		LandingPageURL string `json:"landing_page_url,omitempty"`
		License        string `json:"license,omitempty"`
		Source         struct {
			DisplayName             string   `json:"display_name,omitempty"`
			HostOrganization        string   `json:"host_organization,omitempty"`
			HostOrganizationLineage []string `json:"host_organization_lineage,omitempty"`
			HostOrganizationName    string   `json:"host_organization_name,omitempty"`
			ID                      string   `json:"id,omitempty"`
			ISSNL                   string   `json:"issn_l,omitempty"`
			ISSN                    []string `json:"issn,omitempty"`
			Type                    string   `json:"type,omitempty"`
		} `json:"source,omitempty"`
		PDFURL  string `json:"pdf_url,omitempty"`
		Version string `json:"version,omitempty"`
	} `json:"best_oa_location,omitempty"`
	Biblio struct {
		Volume    string `json:"volume,omitempty"`
		Issue     string `json:"issue,omitempty"`
		FirstPage string `json:"first_page,omitempty"`
		LastPage  string `json:"last_page,omitempty"`
	} `json:"biblio,omitempty"`
	CitedByAPIURL         string `json:"cited_by_api_url,omitempty"`
	CitedByCount          int    `json:"cited_by_count,omitempty"`
	CitedByPercentileYear struct {
		Min float64 `json:"min"`
		Max float64 `json:"max"`
	} `json:"cited_by_percentile_year"`
	Concepts []struct {
		ID          string  `json:"ID,omitempty"`
		Wikidata    string  `json:"wikidata,omitempty"`
		DisplayName string  `json:"display_name,omitempty"`
		Level       int     `json:"level,omitempty"`
		Score       float32 `json:"score,omitempty"`
	} `json:"concepts,omitempty"`
	CorrespondingAuthorIDs      []string `json:"corresponding_author_ids,omitempty"`
	CorrespondingInstitutionIDs []string `json:"corresponding_institution_ids,omitempty"`
	CountriesDistinctCount      int      `json:"countries_distinct_count,omitempty"`
	CountsByYear                []struct {
		Year         int `json:"year,omitempty"`
		CitedByCount int `json:"cited_by_count,omitempty"`
	} `json:"counts_by_year,omitempty"`
	CreatedDate    string `json:"created_date,omitempty"`
	DisplayName    string `json:"display_name,omitempty"`
	DOI            string `json:"doi,omitempty"`
	FullTextOrigin string `json:"fulltext_origin,omitempty"`
	Grants         []struct {
		Funder            string `json:"funder,omitempty"`
		FunderDisplayName string `json:"funder_display_name,omitempty"`
		AwardID           string `json:"award_id,omitempty"`
	} `json:"grants,omitempty"`
	HasFullText bool   `json:"has_fulltext,omitempty"`
	ID          string `json:"id,omitempty"`
	IDs         struct {
		OpenAlex string `json:"openalex,omitempty"`
		DOI      string `json:"doi,omitempty"`
		Mag      string `json:"mag,omitempty"`
		PMID     string `json:"pmid,omitempty"`
	} `json:"ids,omitempty"`
	InstitutionsDistinctCount int  `json:"institutions_distinct_count,omitempty"`
	IsParatext                bool `json:"is_paratext,omitempty"`
	IsRetracted               bool `json:"is_retracted,omitempty"`
	Keywords                  []struct {
		Keyword string  `json:"keyword,omitempty"`
		Score   float32 `json:"score,omitempty"`
	} `json:"keywords,omitempty"`
	Language  string `json:"language,omitempty"`
	License   string `json:"license,omitempty"`
	Locations []struct {
		IsAccepted     bool   `json:"is_accepted,omitempty"`
		IsOA           bool   `json:"is_oa,omitempty"`
		IsPublished    bool   `json:"is_published,omitempty"`
		LandingPageURL string `json:"landing_page_url,omitempty"`
		License        string `json:"license,omitempty"`
		Source         struct {
			DisplayName             string   `json:"display_name,omitempty"`
			HostOrganization        string   `json:"host_organization,omitempty"`
			HostOrganizationLineage []string `json:"host_organization_lineage,omitempty"`
			HostOrganizationName    string   `json:"host_organization_name,omitempty"`
			ID                      string   `json:"id,omitempty"`
			ISSNL                   string   `json:"issn_l,omitempty"`
			ISSN                    []string `json:"issn,omitempty"`
			Type                    string   `json:"type,omitempty"`
		} `json:"source,omitempty"`
		PDFURL  string `json:"pdf_url,omitempty"`
		Version string `json:"version,omitempty"`
	} `json:"locations,omitempty"`
	LocationsCount int `json:"locations_count,omitempty"`
	Mesh           []struct {
		DescriptorUI   string `json:"descriptor_ui,omitempty"`
		DescriptorName string `json:"descriptor_name,omitempty"`
		QualifierUI    string `json:"qualifier_ui,omitempty"`
		QualifierName  string `json:"qualifier_name,omitempty"`
		IsMajorTopic   bool   `json:"is_major_topic,omitempty"`
	} `json:"mesh,omitempty"`
	NGramsURL  string `json:"ngrams_url,omitempty"`
	OpenAccess struct {
		IsOA                     bool   `json:"is_oa,omitempty"`
		OAStatus                 string `json:"oa_status,omitempty"`
		OAURL                    string `json:"oa_url,omitempty"`
		AnyRepositoryHasFullText bool   `json:"any_repository_has_fulltext,omitempty"`
	} `json:"open_access,omitempty"`
	PrimaryLocation struct {
		IsAccepted     bool   `json:"is_accepted,omitempty"`
		IsOA           bool   `json:"is_oa,omitempty"`
		IsPublished    bool   `json:"is_published,omitempty"`
		LandingPageURL string `json:"landing_page_url,omitempty"`
		License        string `json:"license,omitempty"`
		Source         struct {
			DisplayName             string   `json:"display_name,omitempty"`
			HostOrganization        string   `json:"host_organization,omitempty"`
			HostOrganizationLineage []string `json:"host_organization_lineage,omitempty"`
			HostOrganizationName    string   `json:"host_organization_name,omitempty"`
			ID                      string   `json:"id,omitempty"`
			ISSNL                   string   `json:"issn_l,omitempty"`
			ISSN                    []string `json:"issn,omitempty"`
			Type                    string   `json:"type,omitempty"`
		} `json:"source,omitempty"`
		PDFURL  string `json:"pdf_url,omitempty"`
		Version string `json:"version,omitempty"`
	} `json:"primary_location,omitempty"`
	PublicationDate             string   `json:"publication_date,omitempty"`
	PublicationYear             int      `json:"publication_year,omitempty"`
	ReferencedWorks             []string `json:"referenced_works,omitempty"`
	ReferencedWorksCount        int      `json:"referenced_works_count,omitempty"`
	RelatedWorks                []string `json:"related_works,omitempty"`
	SustainableDevelopmentGoals []struct {
		ID          string  `json:"id,omitempty"`
		DisplayName string  `json:"display_name,omitempty"`
		Score       float32 `json:"score,omitempty"`
	} `json:"sustainable_development_goals,omitempty"`
	Title        string `json:"title,omitempty"`
	Type         string `json:"type,omitempty"`
	TypeCrossRef string `json:"type_crossref,omitempty"`
	UpdatedDate  string `json:"updated_date,omitempty"`
}

type Works struct {
	Meta struct {
		Count            float64 `json:"count,omitempty"`
		DBResponseTimeMS int     `json:"db_response_time_ms,omitempty"`
		Page             int     `json:"page,omitempty"`
		PerPage          int     `json:"per_page,omitempty"`
	} `json:"meta,omitempty"`
	Results []Work `json:"results,omitempty"`
}
