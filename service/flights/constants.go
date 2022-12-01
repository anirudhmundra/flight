package flights

const (
	OFFER20 = "OFFER_20"
	OFFER25 = "OFFER_25"
	OFFER30 = "OFFER_30"
)

var getDiscountCodes = struct {
	list map[string]string
}{
	list: map[string]string{
		"A": OFFER20,
		"B": OFFER20,
		"C": OFFER20,
		"D": OFFER20,
		"E": OFFER20,
		"F": OFFER30,
		"G": OFFER30,
		"H": OFFER30,
		"I": OFFER30,
		"J": OFFER30,
		"K": OFFER30,
		"L": OFFER25,
		"M": OFFER25,
		"N": OFFER25,
		"O": OFFER25,
		"P": OFFER25,
		"Q": OFFER25,
		"R": OFFER25,
	},
}
