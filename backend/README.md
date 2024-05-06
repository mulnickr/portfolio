Golang Backend

##Building

<code>
    docker build -t portfolio:alpha .
</code

##Push to Google Artifact Registry

<code>
    docker tag portfolio:alpha gcr.io/rmulnick-web/portfolio:alpha
</code>

<code>
    docker push gcr.io/rmulnick-web/portfolio:alpha
</code>
