package events

import (
    "log"
    "encoding/json"
)

type fruitBasket struct {
    Data string
}

func ToJSON(name string) string {

    basket := fruitBasket{
        Data: name,
    }

    var jsonData []byte
        jsonData, err := json.Marshal(basket)
        if err != nil {
            log.Println(err)
        }
    return string(jsonData)
}
