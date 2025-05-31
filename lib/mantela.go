package mantela_fetcher

// Mantela は Mantela -- Telephone Network Mandala のトップレベルの構造体です。
// 交換局の相互接続情報を提供します。
type Mantela struct {
	// Mantela のバージョンを指定します（通常は '0.0.0' を指定してください）
	Version *string `json:"version"`
	// あなた自身の交換局の情報を提供してください
	AboutMe *AboutMe `json:"aboutMe"`
	// あなた自身の交換局に接続されている端末の情報を提供してください
	Extensions []Extension `json:"extensions"` // 配列自体はnullになり得ますが、要素はnilポインタ
	// 接続されている他の交換局の情報を提供してください
	Providers []Provider `json:"providers"` // 配列自体はnullになり得ますが、要素はnilポインタ
}

// AboutMe はあなた自身の交換局の情報を提供します。
type AboutMe struct {
	// 交換局の名称を人間の読める形式で指定してください
	Name *string `json:"name"`
	// 交換局の好ましいプレフィクス番号を指定してください
	PreferredPrefix interface{} `json:"preferredPrefix"` // string または []string, ポインタ型にすると複雑になるためinterface{}のまま
	// 交換局の識別子を機械の読める形式で指定してください
	Identifier *string `json:"identifier"`
	// 交換局の SIP ユーザ名を指定してください
	SIPUsername *string `json:"sipUsername"`
	// 交換局の SIP パスワードを指定してください
	SIPPassword *string `json:"sipPassword"`
	// 交換局の SIP ホスト名や IP アドレスを指定してください
	SIPServer *string `json:"sipServer"`
	// 交換局の SIP ポート名やポート番号を指定してください
	SIPPort *string `json:"sipPort"`
	// 交換局が（一時的に）利用できないとき、true を指定してください
	Unavailable *bool `json:"unavailable"`
}

// Extension はあなた自身の交換局に接続されている端末の情報を提供します。
// 配列の要素であるため、構造体自体がnilになることはあまり一般的ではありませんが、
// フィールド内の個々の値はnilになり得ます。
type Extension struct {
	// 端末の名称を人間の読める形式で指定してください
	Name *string `json:"name"`
	// 端末に紐付けられる内線番号を指定してください
	Extension *string `json:"extension"`
	// 端末の識別子を機械の読める形式で指定してください
	Identifier *string `json:"identifier"`
	// 端末の種別を指定してください
	Type *string `json:"type"`
	// 転送を伴うような内線番号において、転送先の端末の識別子を指定してください
	TransferTo []string `json:"transferTo"` // 配列自体はnullになり得ますが、要素はnilポインタ
}

// Provider は接続されている他の交換局の情報を提供します。
// 配列の要素であるため、構造体自体がnilになることはあまり一般的ではありませんが、
// フィールド内の個々の値はnilになり得ます。
type Provider struct {
	// 交換局の名称を人間の読める形式で指定してください
	Name *string `json:"name"`
	// 交換局に紐付けられるプレフィクス番号を指定してください
	Prefix *string `json:"prefix"`
	// 交換局の識別子を機械の読める形式で指定してください
	Identifier *string `json:"identifier"`
	// その交換局の mantela.json の URL を指定してください
	Mantela *string `json:"mantela"`
}
