package pocketlog_test

import "pocketlog"

func ExampleLogger_Debugf() {
    debugLogger := pocketlog.New(pocketlog.LevelDebug)     
    debugLogger.Debugf("Hello, %s", "world")
    // Output: Hello, world
}
