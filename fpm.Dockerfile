FROM debian:8
RUN apt-get update && apt-get install -y ruby ruby-dev build-essential
RUN gem install cabin -v 0.9.0
RUN gem install backports -v 3.11.0
RUN gem install arr-pm -v 0.0.10
RUN gem install clamp -v 1.0.1
RUN gem install ffi -v 1.9.18
RUN gem install childprocess -v 0.8.0
RUN gem install io-like -v 0.3.0
RUN gem install ruby-xz -v 0.2.3
RUN gem install dotenv -v 2.2.1
RUN gem install insist -v 1.0.0
RUN gem install mustache -v 0.99.8
RUN gem install stud -v 0.0.23
RUN gem install pleaserun -v 0.0.30
RUN gem install fpm -v 1.9.3
WORKDIR /packaging
