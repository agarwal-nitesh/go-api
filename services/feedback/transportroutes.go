package feedback


type TransportRoutes struct {
	GetFeedbackOptionsAPI HTTPTransportInterface
	SubmitFeedbackAPI     HTTPTransportInterface
}

func MakeTransportRoutes(service FeedbackServiceInterface) (TransportRoutes, error) {
	defaultHTTPTransport := DefaultHTTPTransport{service}
	return TransportRoutes{
		GetFeedbackOptionsAPI: &GetFeedbackOptionsAPI{defaultHTTPTransport},
		SubmitFeedbackAPI:     &CreateSubmitFeedbackAPI{defaultHTTPTransport},
	}, nil
}