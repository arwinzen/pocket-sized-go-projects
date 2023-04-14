package main

import "testing"

// tests output to stdout
// by using Example<Function> Go will test stdout
// against the next line which follows `// Output:`
func ExampleMain() {
    main()
    // Output:
    // Hello world
}

func TestGreet(t *testing.T) {
    // struct to define a test case
    type testCase struct{
        lang language
        expectedGreeting string
    }

    // map to store test cases
    tests := map[string]testCase{
        "English": {
            // we can define the properties without keys
            // since the order is defined as lang then greeting
            "en",
            "Hello world",
        },
        "French": {
            // or we can be explicit by specifying the keys
            lang:             "fr",
            expectedGreeting: "Bonjour le monde",
        },
        "Akkadian, not supported": {
            lang:             "akk",
            expectedGreeting: `unsupported language: "akk"`,
        },
        "Greek": {
            lang:             "el",
            expectedGreeting: "Χαίρετε Κόσμε",
        },
        "Hebrew": {
            lang:             "he",
            expectedGreeting: "שלום עולם",
        },
        "Urdu": {
            lang:             "ur",
            expectedGreeting: "ہیلو دنیا",
        },
        "Vietnamese": {
            lang:             "vi",
            expectedGreeting: "Xin chào Thế Giới",
        },
        "Empty": {
            lang:             "",
            expectedGreeting: `unsupported language: ""`,
        },
    }

    // range over all test cases
    for name, tc := range tests {
        // t.Run runs a closure on a subset of the tests
        // and reports if the test is successful
        t.Run(name, func(t *testing.T) {
            greeting := greet(tc.lang)
            
            if greeting != tc.expectedGreeting {
                t.Errorf("wanted: %q, got: %q", tc.expectedGreeting, greeting)
            }
        })
    } 
}
