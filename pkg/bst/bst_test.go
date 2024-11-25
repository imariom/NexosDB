package bst

import (
	"testing"
	"time"

	kv "github.com/imariom/nexusdb/pkg/kvpair"
)

func TestBST_InsertAndSearch(t *testing.T) {
	kvpairs := []struct {
		key    []byte
		value  []byte
		ttl    time.Duration
		result []byte
	}{
		{[]byte("userID123"), []byte("John Doe"), time.Second * 1, []byte("John Doe")},
		{[]byte("sessionToken"), []byte("abc123xyz"), time.Minute * 5, []byte("abc123xyz")},
		{[]byte("permanentUserID"), []byte("user123456"), 0, []byte("user123456")},
		{[]byte("email"), []byte("jane.doe@example.com"), time.Hour * 1, []byte("jane.doe@example.com")},
		{[]byte("orderID456"), []byte("Order#789456"), time.Minute * 30, []byte("Order#789456")},
		{[]byte("configSetting"), []byte("default"), 0, []byte("default")},
		{[]byte("productID"), []byte("Widget-X100"), time.Hour * 2, []byte("Widget-X100")},
		{[]byte("binaryData"), []byte{0x0A, 0x1B, 0x2C, 0x3D}, time.Second * 45, []byte{0x0A, 0x1B, 0x2C, 0x3D}},
		{[]byte("cartID"), []byte("CART98765"), time.Minute * 15, []byte("CART98765")},
		{[]byte("apiKey"), []byte("apikey-xyz-123"), 0, []byte("apikey-xyz-123")},
		{[]byte("authToken"), []byte("authToken456"), time.Minute * 20, []byte("authToken456")},
		{[]byte("licenseKey"), []byte("license-789-abcd"), 0, []byte("license-789-abcd")},
		{[]byte("userPreference"), []byte("theme:dark"), time.Hour * 4, []byte("theme:dark")},
		{[]byte("orderStatus"), []byte("pending"), time.Hour * 6, []byte("pending")},
		{[]byte("imageHeader"), []byte{0xFF, 0xD8, 0xFF, 0xE0}, time.Minute * 20, []byte{0xFF, 0xD8, 0xFF, 0xE0}},
		{[]byte("promotionCode"), []byte("DISCOUNT20"), time.Minute * 45, []byte("DISCOUNT20")},
		{[]byte("featureToggle"), []byte("enabled"), 0, []byte("enabled")},
		{[]byte("sessionID"), []byte("xyz12345session"), time.Minute * 10, []byte("xyz12345session")},
		{[]byte("jsonString"), []byte("{\"name\":\"John\"}"), time.Hour * 1, []byte("{\"name\":\"John\"}")},
		{[]byte("largeText"), []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit."), time.Hour * 2, []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")},
		{[]byte("integerValue"), []byte("123456789"), 0, []byte("123456789")},
		{[]byte("hexData"), []byte("4A6F686E446F65"), time.Minute * 50, []byte("4A6F686E446F65")},
		{[]byte("orderDetails"), []byte("Order123|Qty:10|Price:$99.99"), time.Hour * 3, []byte("Order123|Qty:10|Price:$99.99")},
		{[]byte("floatValue"), []byte("123.456"), 0, []byte("123.456")},
		{[]byte("binaryBlob"), []byte{0x01, 0x02, 0x03, 0x04, 0x05}, time.Minute * 15, []byte{0x01, 0x02, 0x03, 0x04, 0x05}},
		{[]byte("encodedString"), []byte("U29tZSBlbmNvZGVkIHN0cmluZw=="), time.Hour * 5, []byte("U29tZSBlbmNvZGVkIHN0cmluZw==")},
		{[]byte("customFlag"), []byte("true"), 0, []byte("true")},
		{[]byte("timestamp"), []byte("2024-11-15T15:34:00Z"), time.Hour * 24, []byte("2024-11-15T15:34:00Z")},
		{[]byte("nestedJSON"), []byte("{\"user\":{\"id\":1,\"name\":\"Jane\"}}"), time.Minute * 40, []byte("{\"user\":{\"id\":1,\"name\":\"Jane\"}}")},
		{[]byte("imageData"), []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46}, time.Minute * 20, []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46}},
		{[]byte("booleanValue"), []byte("false"), 0, []byte("false")},
		{[]byte("base64Image"), []byte("/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDABALDA4M"), time.Minute * 60, []byte("/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDABALDA4M")},
		{[]byte("hashValue"), []byte("5d41402abc4b2a76b9719d911017c592"), time.Hour * 12, []byte("5d41402abc4b2a76b9719d911017c592")},
		{[]byte("uuid"), []byte("550e8400-e29b-41d4-a716-446655440000"), 0, []byte("550e8400-e29b-41d4-a716-446655440000")},
		{[]byte("ipv4Address"), []byte("192.168.1.1"), time.Hour * 6, []byte("192.168.1.1")},
		{[]byte("largeBinary"), []byte{0xFF, 0xFE, 0xFD, 0xFC}, time.Minute * 25, []byte{0xFF, 0xFE, 0xFD, 0xFC}},
		{[]byte("ipv6Address"), []byte("2001:0db8:85a3:0000:0000:8a2e:0370:7334"), time.Hour * 8, []byte("2001:0db8:85a3:0000:0000:8a2e:0370:7334")},
		{[]byte("xmlData"), []byte("<user><id>1</id><name>John</name></user>"), time.Hour * 2, []byte("<user><id>1</id><name>John</name></user>")},
		{[]byte("temperature"), []byte("36.6"), time.Minute * 30, []byte("36.6")},
		{[]byte("audioData"), []byte{0x52, 0x49, 0x46, 0x46, 0x24, 0x08, 0x00, 0x00, 0x57, 0x41, 0x56, 0x45}, time.Minute * 10, []byte{0x52, 0x49, 0x46, 0x46, 0x24, 0x08, 0x00, 0x00, 0x57, 0x41, 0x56, 0x45}},
		{[]byte("dnsRecord"), []byte("example.com"), time.Hour * 4, []byte("example.com")},
		{[]byte("shortHex"), []byte("0x1F"), time.Second * 30, []byte("0x1F")},
		{[]byte("sessionData"), []byte("active"), time.Hour * 7, []byte("active")},
		{[]byte("encryptedPayload"), []byte{0x8A, 0x2F, 0xCD, 0xEE, 0x12, 0x34, 0xAB, 0x90}, time.Minute * 5, []byte{0x8A, 0x2F, 0xCD, 0xEE, 0x12, 0x34, 0xAB, 0x90}},
		{[]byte("quote"), []byte("To be or not to be."), time.Minute * 10, []byte("To be or not to be.")},
		{[]byte("csvRow"), []byte("1,John,Doe,30"), time.Hour * 3, []byte("1,John,Doe,30")},
		{[]byte("shortBinary"), []byte{0xAB, 0xCD}, time.Minute * 10, []byte{0xAB, 0xCD}},
		{[]byte("sensorID"), []byte("SENSOR123"), time.Minute * 15, []byte("SENSOR123")},
		{[]byte("transactionID"), []byte("TXN456789"), time.Hour * 1, []byte("TXN456789")},
		{[]byte("binaryFlag"), []byte{0x01}, time.Minute * 5, []byte{0x01}},
		{[]byte("deviceStatus"), []byte("online"), time.Minute * 5, []byte("online")},
		{[]byte("orderSummary"), []byte("Order#123, Status: Shipped"), time.Hour * 2, []byte("Order#123, Status: Shipped")},
		{[]byte("checksum"), []byte{0xDE, 0xAD, 0xBE, 0xEF}, time.Hour * 1, []byte{0xDE, 0xAD, 0xBE, 0xEF}},
	}

	// Inser all previous pair into the bst
	bst := &BST{}
	for _, test := range kvpairs {
		bst.Insert(kv.NewKVPair(test.key, test.value, test.ttl))
	}

	// Test search for inserted nodes
	for _, test := range kvpairs {
		if !bst.Search(test.key) {
			t.Errorf("Expected to find '%s' in BST", test.key)
		}
	}

	// Test search for a non-existent value
	keys := []string{
		"deviceID456",
		"userRole",
		"lastLoginTimestamp",
		"accessLevel",
		"paymentMethod",
		"shippingAddress",
		"billingAddress",
		"accountStatus",
		"referralCode",
		"subscriptionType",
		"membershipLevel",
		"loginAttempts",
		"securityQuestion",
		"passwordResetToken",
		"twoFactorAuth",
		"profilePictureURL",
		"languagePreference",
		"timezoneSetting",
		"notificationSettings",
		"cartItemCount",
		"wishlistItems",
		"favoriteCategories",
		"purchaseHistory",
		"giftCardBalance",
		"loyaltyPoints",
		"affiliateID",
		"browserSession",
		"apiRateLimit",
		"cacheVersion",
		"featureFlags",
		"systemMetrics",
	}

	for _, key := range keys {
		if bst.Search([]byte(key)) {
			t.Errorf("Expected not to find '%s' in BST", key)
		}
	}
}

func TestBST_UpdateAndSearch(t *testing.T) {
	kvpairs := []struct {
		key         []byte
		value       []byte
		ttl         time.Duration
		updateValue []byte
		result      []byte
	}{
		{[]byte("userID123"), []byte("John Doe"), time.Second * 1, []byte("Jane Smith"), []byte("Jane Smith")},
		{[]byte("sessionToken"), []byte("abc123xyz"), time.Minute * 5, []byte("xyz789abc"), []byte("xyz789abc")},
		{[]byte("permanentUserID"), []byte("user123456"), 0, []byte("user789012"), []byte("user789012")},
		{[]byte("email"), []byte("jane.doe@example.com"), time.Hour * 1, []byte("john.doe@example.com"), []byte("john.doe@example.com")},
		{[]byte("orderID456"), []byte("Order#789456"), time.Minute * 30, []byte("Order#123456"), []byte("Order#123456")},
		{[]byte("configSetting"), []byte("default"), 0, []byte("custom"), []byte("custom")},
		{[]byte("productID"), []byte("Widget-X100"), time.Hour * 2, []byte("Widget-Y200"), []byte("Widget-Y200")},
		{[]byte("binaryData"), []byte{0x0A, 0x1B, 0x2C, 0x3D}, time.Second * 45, []byte{0x4E, 0x5F, 0x6A, 0x7B}, []byte{0x4E, 0x5F, 0x6A, 0x7B}},
		{[]byte("cartID"), []byte("CART98765"), time.Minute * 15, []byte("CART54321"), []byte("CART54321")},
		{[]byte("apiKey"), []byte("apikey-xyz-123"), 0, []byte("apikey-abc-456"), []byte("apikey-abc-456")},
		{[]byte("authToken"), []byte("authToken456"), time.Minute * 20, []byte("authToken789"), []byte("authToken789")},
		{[]byte("licenseKey"), []byte("license-789-abcd"), 0, []byte("license-123-efgh"), []byte("license-123-efgh")},
		{[]byte("userPreference"), []byte("theme:dark"), time.Hour * 4, []byte("theme:light"), []byte("theme:light")},
		{[]byte("orderStatus"), []byte("pending"), time.Hour * 6, []byte("completed"), []byte("completed")},
		{[]byte("imageHeader"), []byte{0xFF, 0xD8, 0xFF, 0xE0}, time.Minute * 20, []byte{0xFA, 0xDE, 0xBE, 0xEF}, []byte{0xFA, 0xDE, 0xBE, 0xEF}},
		{[]byte("promotionCode"), []byte("DISCOUNT20"), time.Minute * 45, []byte("DISCOUNT30"), []byte("DISCOUNT30")},
		{[]byte("featureToggle"), []byte("enabled"), 0, []byte("disabled"), []byte("disabled")},
		{[]byte("sessionID"), []byte("xyz12345session"), time.Minute * 10, []byte("abc98765session"), []byte("abc98765session")},
		{[]byte("jsonString"), []byte("{\"name\":\"John\"}"), time.Hour * 1, []byte("{\"name\":\"Jane\"}"), []byte("{\"name\":\"Jane\"}")},
		{[]byte("largeText"), []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit."), time.Hour * 2, []byte("Ut enim ad minim veniam, quis nostrud exercitation."), []byte("Ut enim ad minim veniam, quis nostrud exercitation.")},
		{[]byte("integerValue"), []byte("123456789"), 0, []byte("987654321"), []byte("987654321")},
		{[]byte("hexData"), []byte("4A6F686E446F65"), time.Minute * 50, []byte("5065746572416E"), []byte("5065746572416E")},
		{[]byte("orderDetails"), []byte("Order123|Qty:10|Price:$99.99"), time.Hour * 3, []byte("Order456|Qty:5|Price:$49.99"), []byte("Order456|Qty:5|Price:$49.99")},
		{[]byte("floatValue"), []byte("123.456"), 0, []byte("654.321"), []byte("654.321")},
		{[]byte("binaryBlob"), []byte{0x01, 0x02, 0x03, 0x04, 0x05}, time.Minute * 15, []byte{0x06, 0x07, 0x08, 0x09, 0x0A}, []byte{0x06, 0x07, 0x08, 0x09, 0x0A}},
		{[]byte("encodedString"), []byte("U29tZSBlbmNvZGVkIHN0cmluZw=="), time.Hour * 5, []byte("TmV3IGVuY29kZWQgc3RyaW5n"), []byte("TmV3IGVuY29kZWQgc3RyaW5n")},
		{[]byte("customFlag"), []byte("true"), 0, []byte("false"), []byte("false")},
		{[]byte("sessionTimeout"), []byte("300"), time.Minute * 5, []byte("600"), []byte("600")},
		{[]byte("username"), []byte("guest123"), time.Hour * 1, []byte("user456"), []byte("user456")},
		{[]byte("errorMessage"), []byte("Invalid credentials"), time.Minute * 10, []byte("Access denied"), []byte("Access denied")},
		{[]byte("phoneNumber"), []byte("+1234567890"), time.Hour * 2, []byte("+0987654321"), []byte("+0987654321")},
		{[]byte("currencyCode"), []byte("USD"), time.Minute * 30, []byte("EUR"), []byte("EUR")},
		{[]byte("apiRateLimit"), []byte("1000 requests/min"), 0, []byte("500 requests/min"), []byte("500 requests/min")},
		{[]byte("country"), []byte("Mozambique"), time.Hour * 24, []byte("South Africa"), []byte("South Africa")},
		{[]byte("fileHash"), []byte("1234abcd"), time.Minute * 15, []byte("5678efgh"), []byte("5678efgh")},
		{[]byte("alertStatus"), []byte("inactive"), time.Hour * 4, []byte("active"), []byte("active")},
		{[]byte("discountCode"), []byte("SUMMER2024"), time.Minute * 50, []byte("WINTER2024"), []byte("WINTER2024")},
		{[]byte("featureName"), []byte("BetaFeature1"), 0, []byte("BetaFeature2"), []byte("BetaFeature2")},
		{[]byte("taskID"), []byte("task-001"), time.Minute * 25, []byte("task-002"), []byte("task-002")},
		{[]byte("xmlPayload"), []byte("<name>John</name>"), time.Hour * 1, []byte("<name>Jane</name>"), []byte("<name>Jane</name>")},
		{[]byte("accountBalance"), []byte("1000.00"), time.Hour * 3, []byte("2000.00"), []byte("2000.00")},
		{[]byte("geoLocation"), []byte("25.9692° S, 32.5732° E"), time.Hour * 6, []byte("34.0522° N, 118.2437° W"), []byte("34.0522° N, 118.2437° W")},
		{[]byte("apiVersion"), []byte("v1.0.0"), time.Minute * 20, []byte("v2.0.0"), []byte("v2.0.0")},
		{[]byte("jsonData"), []byte("{\"age\":25}"), time.Hour * 2, []byte("{\"age\":30}"), []byte("{\"age\":30}")},
		{[]byte("objectID"), []byte("obj-0001"), time.Minute * 45, []byte("obj-0002"), []byte("obj-0002")},
		{[]byte("productCategory"), []byte("Electronics"), time.Hour * 1, []byte("Home Appliances"), []byte("Home Appliances")},
		{[]byte("booleanFlag"), []byte("true"), time.Minute * 35, []byte("false"), []byte("false")},
		{[]byte("configID"), []byte("configA"), 0, []byte("configB"), []byte("configB")},
		{[]byte("requestToken"), []byte("req-12345"), time.Minute * 40, []byte("req-67890"), []byte("req-67890")},
		{[]byte("binaryPayload"), []byte{0x10, 0x20, 0x30, 0x40}, time.Hour * 5, []byte{0x50, 0x60, 0x70, 0x80}, []byte{0x50, 0x60, 0x70, 0x80}},
		{[]byte("systemMode"), []byte("idle"), time.Minute * 30, []byte("active"), []byte("active")},
		{[]byte("retryCount"), []byte("3"), 0, []byte("5"), []byte("5")},
		{[]byte("hexString"), []byte("ABCDEF"), time.Minute * 25, []byte("123456"), []byte("123456")},
		{[]byte("apiKeyPrefix"), []byte("KEY-"), time.Hour * 6, []byte("SECRET-"), []byte("SECRET-")},
		{[]byte("dataChecksum"), []byte("A1B2C3"), time.Minute * 15, []byte("D4E5F6"), []byte("D4E5F6")},
		{[]byte("tempValue"), []byte("25.5°C"), time.Minute * 10, []byte("30.0°C"), []byte("30.0°C")},
	}

	// Inser all previous pair into the bst
	bst := &BST{}
	for _, test := range kvpairs {
		bst.Insert(kv.NewKVPair(test.key, test.value, test.ttl))
	}

	// Update all previous inserted nodes
	for _, test := range kvpairs {
		bst.Insert(kv.NewKVPair(test.key, test.updateValue, test.ttl))
	}

	// Get the updated KVPair corresponding for each key
	for _, test := range kvpairs {
		node, err := bst.Get(test.key)
		if err != nil {
			t.Fatalf("Expected KVPair, but got %v", err)
		}

		// Get the value from the previous KVPair
		value, err := node.Value()
		if err != nil || !slicesEqual(value, test.result) {
			t.Errorf("Expected update '%s' value, but got %v", string(test.result), err)
		}
	}
}

func TestBST_InOrderTraversal(t *testing.T) {
}

func slicesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestBST_GetMethod(t *testing.T) {
	kvpairs := []struct {
		key    []byte
		value  []byte
		ttl    time.Duration
		result []byte
	}{
		{[]byte("userID123"), []byte("John Doe"), time.Second * 1, []byte("John Doe")},
		{[]byte("sessionToken"), []byte("abc123xyz"), time.Minute * 5, []byte("abc123xyz")},
		{[]byte("permanentUserID"), []byte("user123456"), 0, []byte("user123456")},
		{[]byte("email"), []byte("jane.doe@example.com"), time.Hour * 1, []byte("jane.doe@example.com")},
		{[]byte("orderID456"), []byte("Order#789456"), time.Minute * 30, []byte("Order#789456")},
		{[]byte("configSetting"), []byte("default"), 0, []byte("default")},
		{[]byte("productID"), []byte("Widget-X100"), time.Hour * 2, []byte("Widget-X100")},
		{[]byte("binaryData"), []byte{0x0A, 0x1B, 0x2C, 0x3D}, time.Second * 45, []byte{0x0A, 0x1B, 0x2C, 0x3D}},
		{[]byte("cartID"), []byte("CART98765"), time.Minute * 15, []byte("CART98765")},
		{[]byte("apiKey"), []byte("apikey-xyz-123"), 0, []byte("apikey-xyz-123")},
		{[]byte("authToken"), []byte("authToken456"), time.Minute * 20, []byte("authToken456")},
		{[]byte("licenseKey"), []byte("license-789-abcd"), 0, []byte("license-789-abcd")},
		{[]byte("userPreference"), []byte("theme:dark"), time.Hour * 4, []byte("theme:dark")},
		{[]byte("orderStatus"), []byte("pending"), time.Hour * 6, []byte("pending")},
		{[]byte("imageHeader"), []byte{0xFF, 0xD8, 0xFF, 0xE0}, time.Minute * 20, []byte{0xFF, 0xD8, 0xFF, 0xE0}},
		{[]byte("promotionCode"), []byte("DISCOUNT20"), time.Minute * 45, []byte("DISCOUNT20")},
		{[]byte("featureToggle"), []byte("enabled"), 0, []byte("enabled")},
		{[]byte("sessionID"), []byte("xyz12345session"), time.Minute * 10, []byte("xyz12345session")},
		{[]byte("jsonString"), []byte("{\"name\":\"John\"}"), time.Hour * 1, []byte("{\"name\":\"John\"}")},
		{[]byte("largeText"), []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit."), time.Hour * 2, []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")},
		{[]byte("integerValue"), []byte("123456789"), 0, []byte("123456789")},
		{[]byte("hexData"), []byte("4A6F686E446F65"), time.Minute * 50, []byte("4A6F686E446F65")},
		{[]byte("orderDetails"), []byte("Order123|Qty:10|Price:$99.99"), time.Hour * 3, []byte("Order123|Qty:10|Price:$99.99")},
		{[]byte("floatValue"), []byte("123.456"), 0, []byte("123.456")},
		{[]byte("binaryBlob"), []byte{0x01, 0x02, 0x03, 0x04, 0x05}, time.Minute * 15, []byte{0x01, 0x02, 0x03, 0x04, 0x05}},
		{[]byte("encodedString"), []byte("U29tZSBlbmNvZGVkIHN0cmluZw=="), time.Hour * 5, []byte("U29tZSBlbmNvZGVkIHN0cmluZw==")},
		{[]byte("customFlag"), []byte("true"), 0, []byte("true")},
		{[]byte("timestamp"), []byte("2024-11-15T15:34:00Z"), time.Hour * 24, []byte("2024-11-15T15:34:00Z")},
		{[]byte("nestedJSON"), []byte("{\"user\":{\"id\":1,\"name\":\"Jane\"}}"), time.Minute * 40, []byte("{\"user\":{\"id\":1,\"name\":\"Jane\"}}")},
		{[]byte("imageData"), []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46}, time.Minute * 20, []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46}},
		{[]byte("booleanValue"), []byte("false"), 0, []byte("false")},
		{[]byte("base64Image"), []byte("/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDABALDA4M"), time.Minute * 60, []byte("/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDABALDA4M")},
		{[]byte("hashValue"), []byte("5d41402abc4b2a76b9719d911017c592"), time.Hour * 12, []byte("5d41402abc4b2a76b9719d911017c592")},
		{[]byte("uuid"), []byte("550e8400-e29b-41d4-a716-446655440000"), 0, []byte("550e8400-e29b-41d4-a716-446655440000")},
		{[]byte("ipv4Address"), []byte("192.168.1.1"), time.Hour * 6, []byte("192.168.1.1")},
		{[]byte("largeBinary"), []byte{0xFF, 0xFE, 0xFD, 0xFC}, time.Minute * 25, []byte{0xFF, 0xFE, 0xFD, 0xFC}},
		{[]byte("ipv6Address"), []byte("2001:0db8:85a3:0000:0000:8a2e:0370:7334"), time.Hour * 8, []byte("2001:0db8:85a3:0000:0000:8a2e:0370:7334")},
		{[]byte("xmlData"), []byte("<user><id>1</id><name>John</name></user>"), time.Hour * 2, []byte("<user><id>1</id><name>John</name></user>")},
		{[]byte("temperature"), []byte("36.6"), time.Minute * 30, []byte("36.6")},
		{[]byte("audioData"), []byte{0x52, 0x49, 0x46, 0x46, 0x24, 0x08, 0x00, 0x00, 0x57, 0x41, 0x56, 0x45}, time.Minute * 10, []byte{0x52, 0x49, 0x46, 0x46, 0x24, 0x08, 0x00, 0x00, 0x57, 0x41, 0x56, 0x45}},
		{[]byte("dnsRecord"), []byte("example.com"), time.Hour * 4, []byte("example.com")},
		{[]byte("shortHex"), []byte("0x1F"), time.Second * 30, []byte("0x1F")},
		{[]byte("sessionData"), []byte("active"), time.Hour * 7, []byte("active")},
		{[]byte("encryptedPayload"), []byte{0x8A, 0x2F, 0xCD, 0xEE, 0x12, 0x34, 0xAB, 0x90}, time.Minute * 5, []byte{0x8A, 0x2F, 0xCD, 0xEE, 0x12, 0x34, 0xAB, 0x90}},
		{[]byte("quote"), []byte("To be or not to be."), time.Minute * 10, []byte("To be or not to be.")},
		{[]byte("csvRow"), []byte("1,John,Doe,30"), time.Hour * 3, []byte("1,John,Doe,30")},
		{[]byte("shortBinary"), []byte{0xAB, 0xCD}, time.Minute * 10, []byte{0xAB, 0xCD}},
		{[]byte("sensorID"), []byte("SENSOR123"), time.Minute * 15, []byte("SENSOR123")},
		{[]byte("transactionID"), []byte("TXN456789"), time.Hour * 1, []byte("TXN456789")},
		{[]byte("binaryFlag"), []byte{0x01}, time.Minute * 5, []byte{0x01}},
		{[]byte("deviceStatus"), []byte("online"), time.Minute * 5, []byte("online")},
		{[]byte("orderSummary"), []byte("Order#123, Status: Shipped"), time.Hour * 2, []byte("Order#123, Status: Shipped")},
		{[]byte("checksum"), []byte{0xDE, 0xAD, 0xBE, 0xEF}, time.Hour * 1, []byte{0xDE, 0xAD, 0xBE, 0xEF}},
	}

	// Inser all previous pair into the bst
	bst := &BST{}
	for _, test := range kvpairs {
		bst.Insert(kv.NewKVPair(test.key, test.value, test.ttl))
	}

	// Get the KVPair corresponding to current key
	for _, test := range kvpairs {
		node, err := bst.Get(test.key)
		if err != nil {
			t.Fatalf("Expected KVPair, but got %v", err)
		}

		// Get the value from the previous KVPair
		value, err := node.Value()
		if err != nil || !slicesEqual(value, test.result) {
			t.Errorf("Expected '%s' value, but got %v", string(test.result), err)
		}
	}
}

func TestBST_Delete(t *testing.T) {
}
