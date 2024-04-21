# KrakenD Gateway Dockerfile
FROM devopsfaith/krakend:2.6.0

# Set environment variables
ENV INFLUX_HOST=${INFLUX_HOST}
ENV INFLUX_PORT=${INFLUX_PORT}
ENV INFLUX_USER=${INFLUX_USER}
ENV INFLUX_USER_PASSWORD=${INFLUX_USER_PASSWORD}

# Copy the configuration file template
COPY krakend.template.json /etc/krakend/krakend.template.json

# Install gettext
RUN apk add gettext

# Set the entrypoint to run envsubst and start KrakenD based of the template
CMD envsubst < /etc/krakend/krakend.template.json > /etc/krakend/krakend.json && krakend run -c /etc/krakend/krakend.json