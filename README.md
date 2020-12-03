# Recipe API :tomato:

This is an amazing API to get the best recipes by the ingredients. The API provides an endpoint `/repices` where you will send your ingredients list and will receive the recipe response.

    GET {HOST}/recipes?i={INGREDIENT_1},{INGREDIENT_2},{INGREDIENT_3}

### Example

Call this API endpoint
    
    GET http://localhost:8080/recipes?i=sugar,salt,potato

and receive this response:

    {
     "Keywords": [
        "sugar",
        "salt",
        "potato"
      ],
      "Recipes": [
        {
          "Title": "Sugar - Browned Potatoes",
          "Ingredients": [
            "butter",
            "potato",
            "salt",
            "sugar"
          ],
          "Link": "http://www.recipezaar.com/Sugar-Browned-Potatoes-133941",
          "Gif": "https://media4.giphy.com/media/ezffuy3sSkpUc/giphy.gif?cid=f3467094zn49kscgufhpqpnmfmw1l16q704zxgi2d5piy6xo&rid=giphy.gif"
        },
        {
          "Title": "German Saurkraut Recipe",
          "Ingredients": [
            "butter",
            "potato",
            "salt",
            "sugar"
          ],
          "Link": "http://www.grouprecipes.com/56420/german-saurkraut.html",
          "Gif": "https://media0.giphy.com/media/Aj9EHGocwb4bu/giphy.gif?cid=f3467094xmr5nwrp56l6bcmiufgdjk0d8ndz243jz2rc99q4&rid=giphy.gif"
        }
      ]
    }

### API limitation

-   You need to send max three ingredients;
-   All ingredients should be in English.

## How run with Docker?

At the project folder run these two commands.
 *Maybe you should run with **sudo** permission

    # Run this command to build
    sudo docker build -t therecipe .
    
    # Run this command to run the api at port 8080 
    sudo docker run -it -p 8080:8080 therecipe

## How run without Docker?

For you to run this project without docker, you need to have the golang installed on your computer. Clone this project in this path

    $GOPATH/src/github.com/SamuelBFavarin/

 Also, you should change your go variable environment to **off**

    GO111MODULE=off

Run the tests

    go test ./...

After, you just should to go to `./source` path and run this command

    go run main.go

