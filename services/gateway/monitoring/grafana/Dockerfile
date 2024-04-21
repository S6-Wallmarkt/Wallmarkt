# Grafana Dockerfile
FROM grafana/grafana:9.1.2

# Set environment variables
ENV INFLUX_HOST=${INFLUX_HOST}
ENV INFLUX_PORT=${INFLUX_PORT}
ENV INFLUX_USER=${INFLUX_USER}
ENV INFLUX_USER_PASSWORD=${INFLUX_USER_PASSWORD}

# Expose port 3000
EXPOSE 3000

# Copy dashboard template
COPY ./dashboards/all.yml /etc/grafana/provisioning/dashboards/all.yml
COPY ./krakend /var/lib/grafana/dashboards/krakend

# Copy datasource template
COPY ./datasources/all.template.yml /etc/grafana/provisioning/datasources/all.template.yml

# Change the user to root to install gettext
USER root

# Install gettext
RUN apk add --no-cache gettext

# Set the entrypoint to run envsubst when the container starts
ENTRYPOINT ["sh", "-c", "envsubst < /etc/grafana/provisioning/datasources/all.template.yml > /etc/grafana/provisioning/datasources/all.yml && exec grafana-server"]

# Change the user to grafana for security reasons
USER grafana