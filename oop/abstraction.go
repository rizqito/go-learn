package main

import "fmt"

type Notification interface {
	Send(tujuan, isi string)
}

type Email struct {
	sender      string
	destination string
	subject     string
	body        string
}

func (e *Email) Send(destination, isi string) {
	e.destination = destination
	e.body = isi
	fmt.Printf("Kirim EMail ke %s dari %s, dengan subject %s, body: %s.\n", e.destination, e.sender, e.subject, e.body)
}

type SMS struct {
	sender      string
	destination string
	teks        string
}

func (s *SMS) Send(destination, isi string) {
	s.destination = destination
	s.teks = isi
	fmt.Printf("Kirim SMS ke %s dari %s, dengan pesan %s.\n", s.destination, s.sender, s.teks)
}

type NotificationSender struct {
	notifications map[string]Notification
}

type NotifParams struct {
	destination string
	content     string
}

func (np *NotifParams) isValid() bool {
	return np.content != "" && np.destination != ""
}

func (ns *NotificationSender) Send(medium string, params NotifParams) {
	if !params.isValid() {
		return
	}

	notif, ok := ns.notifications[medium]
	if !ok {
		return
	}

	notif.Send(params.destination, params.content)
}

func abstraction() {
	fmt.Println("====abstraction====")
	ns := NotificationSender{
		notifications: map[string]Notification{
			"email": &Email{
				sender:  "me@gmail.com",
				subject: "no-reply",
			},
			"sms": &SMS{
				sender: "082350790354",
			},
		},
	}

	ns.Send("email", NotifParams{
		destination: "tujuan@gmail.com",
		content:     "ini email",
	})

	ns.Send("sms", NotifParams{
		destination: "08123456780",
		content:     "ini sms",
	})
}
