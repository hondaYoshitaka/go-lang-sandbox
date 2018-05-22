FROM ubuntu:18.04

RUN apt-get update && apt-get upgrade

# Enabled to login with ssh
RUN apt-get install -y openssh-server
RUN mkdir /var/run/sshd

RUN echo 'root:root' | chpasswd

RUN sed -ri 's/^#?PermitRootLogin\s+.*/PermitRootLogin yes/' /etc/ssh/sshd_config
RUN sed -ri 's/UsePAM yes/#UsePAM yes/g' /etc/ssh/sshd_config

RUN mkdir /root/.ssh

# Install go-lang
RUN apt-get install -y golang && go version

# Clean up
RUN apt-get clean
RUN rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Start ssh daemon
EXPOSE 22
CMD ["/usr/sbin/sshd", "-D"]
