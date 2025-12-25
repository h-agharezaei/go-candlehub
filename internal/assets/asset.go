package assets

type AssetType string

const (
    AssetGold   AssetType = "gold"
    AssetForex  AssetType = "forex"
    AssetCrypto AssetType = "crypto"
)

type Asset struct {
    Symbol string
    Type   AssetType
}
