package zarinpal

const (
	ZARINPAL_REQUEST_API_ENDPOINT = "https://api.zarinpal.com/pg/v4/payment/request.json"
	ZARINPAL_VERIFY_API_ENDPOINT  = "https://api.zarinpal.com/pg/v4/payment/verify.json"

	ZARINPAL_REQUEST_SANDBOX_API_ENDPOINT = "https://sandbox.zarinpal.com/pg/v4/payment/request.json"
	ZARINPAL_VERIFY_SANDBOX_API_ENDPOINT  = "https://sandbox.zarinpal.com/pg/v4/payment/verify.json"
)

func New(merchantID string, sandbox bool) (*Zarinpal, error) {

}
