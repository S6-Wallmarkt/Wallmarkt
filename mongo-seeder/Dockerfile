FROM mongo

COPY product.products.json /product.products.json

CMD mongoimport --uri mongodb://mongodb:27017/product --collection products --drop --file /product.products.json --jsonArray