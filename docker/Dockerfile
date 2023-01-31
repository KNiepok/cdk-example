FROM public.ecr.aws/spacelift/runner-terraform:latest
USER root

# Install packages
RUN apk update && apk add --update --no-cache \
    python3 \
    py3-pip \
    npm
RUN apk --no-cache add --virtual builds-deps build-base python3
# Update NPM
RUN npm update -g
# Install AWSCLI
RUN pip install --upgrade pip && \
    pip install --upgrade awscli
# Install cdk
RUN npm install -g aws-cdk
RUN cdk --version

# Add Go
COPY --from=golang:1.19-alpine /usr/local/go/ /usr/local/go/

ENV PATH="/usr/local/go/bin:${PATH}"
RUN go version