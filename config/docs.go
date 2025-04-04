// Package config provides an easy way to get all the critical values of the project.
//
// # An comprehensive guide line defined to followed by every developer
//
// Usage:
//
// This is particularly useful when you don't wont to hardcode the values in the project
//
// # It requires you to add the following AWS credentials as env
//
// Response from the package:
// 1. If the key is found it will return the value
// 2. If the key is not found and has default value. It will return the default value
// 3. If the key is not found and has no default value. It will return 'key not found' error.
package config
