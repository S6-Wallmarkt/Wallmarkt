# KrakenD Gateway Dockerfile
FROM devopsfaith/krakend:2.6.0

# Set environment variables
ENV INFLUX_HOST=${INFLUX_HOST}
ENV INFLUX_PORT=${INFLUX_PORT}
ENV INFLUX_USER=${INFLUX_USER}
ENV INFLUX_USER_PASSWORD=${INFLUX_USER_PASSWORD}

ENV PRODUCT_URL=${PRODUCT_URL}
ENV ORDER_URL=${ORDER_URL}
ENV SHIPPING_URL=${SHIPPING_URL}

# Copy the configuration file template
COPY krakend.template.json /etc/krakend/krakend.template.json

# Install gettext
RUN apk --no-cache add gettext

# Create a non-root user
RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot

# Change the ownership of the /etc/krakend directory to the nonroot user
RUN chown -R nonroot:nonroot /etc/krakend

# Set the nonroot user
USER nonroot

EXPOSE 8080 80 443

# Set the entrypoint to run envsubst and start KrakenD based on the template
CMD envsubst < /etc/krakend/krakend.template.json > /etc/krakend/krakend.json && krakend run -c /etc/krakend/krakend.json
