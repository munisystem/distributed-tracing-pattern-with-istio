FROM ruby:2.5.0-alpine

WORKDIR /app

COPY Gemfile Gemfile.lock ./
RUN bundle

COPY server.rb ./

CMD ["bundle", "exec", "ruby", "server.rb", "-o", "0.0.0.0"]
