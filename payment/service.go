package payment

import (
	"bwastartup/user"
	"strconv"

	midtrans "github.com/veritrans/go-midtrans"
)


type service struct {
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *service{
	return &service{}
}

//MIDTRANS
func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error){
//Code dari https://github.com/veritrans/go-midtrans
    midclient := midtrans.NewClient()
    midclient.ServerKey = "a252525"
    midclient.ClientKey = "a252525"
    midclient.APIEnvType = midtrans.Sandbox

    snapGateway := midtrans.SnapGateway{
        Client: midclient,
    }
	//daftarkan transaksi
	snapReq := &midtrans.SnapReq{
			CustomerDetail: &midtrans.CustDetail{
				Email: user.Email,
				FName: user.Name,
			},
			TransactionDetails: midtrans.TransactionDetails{
				OrderID: strconv.Itoa(transaction.ID),
				GrossAmt: int64(transaction.Amount),
			},
	}
	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil{
		return "", err
	}
	//dapatkan redirect url
	return snapTokenResp.RedirectURL, nil
}