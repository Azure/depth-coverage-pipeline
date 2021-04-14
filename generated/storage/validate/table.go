package validate

import (
    "fmt"
    "regexp"
)

func StorageTableName(i interface{}, k string) (warnings []string, errors []error) {
    v, ok := i.(string)
    if !ok {
        errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
        return
    }
    if len(v) < 3 {
        errors = append(errors, fmt.Errorf("length should be greater than %d", 3))
        return
    }
    if len(v) > 63 {
        errors = append(errors, fmt.Errorf("length should be less than %d", 63))
        return
    }
    if !regexp.MustCompile(`^[A-Za-z][A-Za-z0-9]{2,62}$`).MatchString(v) {
        errors = append(errors, fmt.Errorf("expected value of %s not match regular expression, got %v", k, v))
        return
    }
    return
}
