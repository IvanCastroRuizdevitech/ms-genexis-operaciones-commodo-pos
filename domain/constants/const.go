package constants

var ID_POS int64 = 0

const (
	DB_CON    string = "postgres://pos_master:24002A8ñE5E588D09F41E3A230Ñ01F7E8@localhost:9040/pos_transacciones?sslmode=allow"
	HOST_PORT string = "18887"
	HOST_IP   string = "0.0.0.0"
)

// DefaultHeaders centralize the common headers required by external services.
var DefaultHeaders = map[string]string{
	"authorization":            "12246655",
	"aplicacion":               "GENEXIS",
	"uuid":                     "530d4dc7-4ebb-4bce-a661-a6d218a657ef",
	"fecha":                    "2020-08-29T11:43:24-05:00",
	"identificadordispositivo": "localhost",
	"versionapp":               "1",
	"versioncode":              "1",
	"user-agent":               "Java",
	"Content-Type":             "application/json",
	"Accept":                   "*/*",
	"key":                      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbiI6Ik1JSUNXd0lCQUFLQmdGOTliWGgreFNnanZLNkpjdTRvTS8xcUg1dWZkbzFTR3NBLzk5S2VuRS9aVHczTTJQZlVJUisvUjZmblcwMENWeDl1RlJERXk2UnNCYWQ2UVBISDVQazFYbjJ6N2ZleFRhUFVrckhmSUtBbUV2NURBN2JVZHhWR3ZGZlRQd1NkSE9Ncit1OFJzek00NnBrVGxZYWRDYUxoSWxOL0FxMEZ4WGxMY1M0QlQ5dVBBZ01CQUFFQ2dZQlhsMkpYYVZISXBCOXJxc0UySE1LejdpallROG9uait6aDlNdUs4OXRHazN1a21RZGRMY1RFcmwwNGMweVlzMURmajhiYS91b2VCTWlhKzlXZTgxSGc5aEhieFlrR3FxVmliL1BpbXJEZ1kwdmV1YTJpZVpJaERGN05tNHBVOVVjM002ckdDd3RVZFBNby9QT0Q1WVpNSzJTa1VLOUJHNFZ4UDZhV2FXT01BUUpCQUxISTF1dmZ2c1BwVnF3RmFNZXpjcDd2QWhwU2N5aTk5WFJoUWsrT2VQbmFvTlk2N2RxVGMrYm5RbFd6N2gzN2c0bjZIWEVnaWphME0wTzZ5aGZaaUlFQ1FRQ0pnQlExY1Z1Y0FTQmFud1lXUEgvVmk3WTZsK2FFUlR1UEczUmg0TCt4eWJFNnRZV1BTT2NwUEdUQ2tEalFPb1dzQkFRRlA4TkZtQ2tSY0gzNVF0d1BBa0JTeEswQ2JVMVR5aHRXeE1IdVR5WHF4bDBWUWhOcEFWNzN5cHZqaS9IWWFLZCtkVHJ0Q3I3bjNmeHRCL0t3dExxZGp2c01BdGsxOUpqMVRFM3Y0bytCQWtBRXZ1enpvMFA4TUhYRGF2VzAwc1crd3lHNUNBRW0wNVRvclpQb3hPOFkzZE5HS1huQXBrMkVPTmRwZDd6ZVF3YldGeFZrY1B6MmEvZ082QWtSRDJPMUFrRUFvR3dSSlhkc3VtcFExWTB1SXc1M0wxYWltRnhRSW4yZzIwb1kwRTZHM1duOFZ3VlJ1WGlqS3lHejhHbTNBSDl2cjRZd1dwRlFkN05UeERySUl2VDNjdz09IiwiaWF0IjoxNTk0MDQ5NDQwfQ.QkZV9B7DyzMW52m1XIzxIwOTgfF7NqAkXLfrY6MPJ3g",
	"centroSoporte":            "D3VIT3CH",
}
