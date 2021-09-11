package controller

// https://dksshddl.tistory.com/entry/GO-https-%EC%A0%9C%EA%B3%B5%EC%9D%84-%EC%9C%84%ED%95%9C-%EC%9D%B8%EC%A6%9D%EB%90%9C-SSL%EA%B3%BC-%EC%84%9C%EB%B2%84-%EA%B0%9C%EC%9D%B8-%ED%82%A4-%EC%83%9D%EC%84%B1%ED%95%98%EA%B8%B0

// max := new(big.Int).Lsh(big.NewInt(1), 128)
// serialNumber, _ := rand.Int(rand.Reader, max)

// subject := pkix.Name{
// 	Organization:       []string{"test Organization"},
// 	OrganizationalUnit: []string{"test"},
// 	CommonName:         "Go Web Programming",
// }

// template := x509.Certificate{
// 	SerialNumber: serialNumber,
// 	Subject:      subject,
// 	NotBefore:    time.Now(),
// 	NotAfter:     time.Now().Add(365 * 24 * time.Hour),
// 	KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
// 	ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
// 	IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
// }

// pk, _ := rsa.GenerateKey(rand.Reader, 2048)

// derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
// certOut, _ := os.Create("cert.pem")
// pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
// certOut.Close()

// keyOut, _ := os.Create("key.pem")
// pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
// keyOut.Close()
