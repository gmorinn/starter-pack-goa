package api

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func firebaseClient(ctx context.Context) (*firebase.App, error) {
	config := &firebase.Config{ProjectID: os.Getenv("FIREBASE_PROJECT_ID")}
	client, err := firebase.NewApp(ctx, config, option.WithCredentialsJSON([]byte(`{
		"type": "service_account",
		"project_id": "starterpack-85abc",
		"private_key_id": "26bed26d6229f7596ffa666c76ff14693f28eef8",
		"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC2ta46WHQwn5nO\nAXBLObDELwpGQKwP4A53ajJGJn+cksrHPQVoC1bPYOVOSsHLAnUv1YZxGiX0Sa+T\n7Jb9vzJXosRct4YFRCj0uy9LnTNNO4sVNnhbX5qsdDfVJ4ezczbsdhGx1LGpOHII\njFz31UlsCX4HiJD9LzAvfkSDeEY62BZlMFb9oX/N0QoVhmYmptogZBdCjTf9yK43\noNkOveMiXxzU0rfSTMZNTUTJPx909/IMnp77JDvKaLuYRNw7z0hhAuC6/4GNHvHe\ni+OjMH11GQssQwHIoX5o6vx4Y/vdLU2v+dpP2FN3V9ggDdARkHS2s9aCklZUeVIN\nfbBX+azlAgMBAAECggEAKQeBeY61gNnGUzGVgv8Qh+28U+xf9yyscvqBEanBjGqW\naMl2NlgdhXJP7BXzkOAbasN7blid307Qw9okShUue+mMSPnZecNFozeJtmQ1ygku\nAXR68KH2fPWYOCCpY43i7pTRAmst/j8ztGG42tqlo0I6b/yl+oUnP/MN9awDsWDF\naatexFYxsykz5BiO2xNxXYsxPrVJXHHSyzxzynoWb4jqG/4uaz5prhb9GH7rsJqd\nZTGil/B198SRynNh82VxgUEGIBRYcJhNFJOB8uexAg6fb4mhf6AydlHimPJFYc+N\nE1kAmo77XoEY2+mzKsnod4aouzFQUoyjvuZXtxQZ3QKBgQDxFTJPQj71/B+J8q2m\nYeTtg+uyyInHdRNfoWHFQqLQ8lmFPhOSJ0Yv5lHyw6JLtq1t46yOw/uUgLC5EOVp\nQJrLUnJgq71ZRJDN/GsN46nP2BzFrvH6om0SRU1M5iYwH5Ds+XpJszMTHig1XePt\ngiTHSKpbGoeJKB3EK6yVgRj+zwKBgQDCA9dn8q8OwwOIoxLdupf+uDAe3TlzCyYk\nqg4KcKnDNEKVrVHz14HdutSl/u7FYimj1tvWH8MEWK+4X8vqnAfss8n3Wgfbsecn\nv0I9x5GaiRMjGNwvo9Xum+C+kYxCIOtRTYTfBKOeFK1TG3WEM7sma70dO7gmShp+\nCxxTrawmCwKBgQCWwWCCe7SPcifuxZidUN4tQK9lT04JMkb5KBvRlMfnzYT5c0Yd\nwP18+pQIMRoOKp/s9dB0Pk3WIVthfxVWahXiSx5HIV34so2yocP4HYOU+1mSrMGY\nr3VxVXDvQ4Rh13tzFNbJu4uWyubt6FadsSJ6lBAjhEV5GWLzxzPXEbGltQJ/Oxzw\nhLtsmxLQANGu7d/sb6+u+bhtyJ+c/47QuF2bW96Qawpj7ee7qdNC91C1b255QXI1\niQgkYy6iW4rUlqbzwysEXlkQC+AojfO4OxqTQoKBUjBA5kME666Sy32NPpjJGfyE\nf27qaz/hj3DOXXlljE9HT6DiesgZSk4YU2aXDQKBgQDeCA9KevASuaBtqQQCJgKX\n5aP9jExDvAZJUwtBuGGkT1bP9UXiNifW5jeTLAi3xQrBLytivuxEf0gksCubBX8t\nsboD+QR5PSRkCszTAI1OTSwgXNRFr2aQbLR+D4o64ZT5Jvb9WBrY4aKWRY0fHBrP\n+c1c4Qq0K9B5zLaXsL5lKg==\n-----END PRIVATE KEY-----\n",
		"client_email": "firebase-adminsdk-3kzq8@starterpack-85abc.iam.gserviceaccount.com",
		"client_id": "101146630415913624383",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-3kzq8%40starterpack-85abc.iam.gserviceaccount.com"
	  }
	`)))
	if err != nil {
		return &firebase.App{}, fmt.Errorf("FIREBASE_NEWAPP %s", err.Error())
	}
	return client, nil
}
