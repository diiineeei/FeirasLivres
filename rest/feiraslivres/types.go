package feiraslivres

type FeiraLivreInsert struct {
	LONGITUDE  int    `json:"longitude"`
	LATITUDE   int    `json:"latitude"`
	SETCENS    int    `json:"setcens"`
	AREAP      int    `json:"areap"`
	CODDIST    int    `json:"coddist"`
	DISTRITO   string `json:"distrito"`
	CODSUBPREF int    `json:"codsubpref"`
	SUBPREFE   string `json:"subprefe"`
	REGIAO5    string `json:"regiao5"`
	REGIAO8    string `json:"regiao8"`
	NOME_FEIRA string `json:"nome_feira"`
	REGISTRO   string `json:"registro"`
	LOGRADOURO string `json:"logradouro"`
	NUMERO     string `json:"numero"`
	BAIRRO     string `json:"bairro"`
	REFERENCIA string `json:"referencia"`
}

type FeiraLivre struct {
	ID         int    `json:"id"`
	LONGITUDE  int    `json:"longitude"`
	LATITUDE   int    `json:"latitude"`
	SETCENS    int    `json:"setcens"`
	AREAP      int    `json:"areap"`
	CODDIST    int    `json:"coddist"`
	DISTRITO   string `json:"distrito"`
	CODSUBPREF int    `json:"codsubpref"`
	SUBPREFE   string `json:"subprefe"`
	REGIAO5    string `json:"regiao5"`
	REGIAO8    string `json:"regiao8"`
	NOME_FEIRA string `json:"nome_feira"`
	REGISTRO   string `json:"registro"`
	LOGRADOURO string `json:"logradouro"`
	NUMERO     string `json:"numero"`
	BAIRRO     string `json:"bairro"`
	REFERENCIA string `json:"referencia"`
}

type IDFeiraLivre struct {
	ID int `json:"id"`
}
