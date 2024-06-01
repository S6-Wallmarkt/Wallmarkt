FROM mongo

COPY product.products.json /product.products.json
COPY order.orders.json /order.orders.json
COPY shipping.shipments.json /shipping.shipments.json

CMD mongoimport --uri mongodb://mongodb:27017/product --collection products --drop --file /product.products.json --jsonArray && mongoimport --uri mongodb://mongodb:27017/order --collection orders --drop --file /order.orders.json --jsonArray && mongoimport --uri mongodb://mongodb:27017/shipping --collection shipments --drop --file /shipping.shipments.json --jsonArray