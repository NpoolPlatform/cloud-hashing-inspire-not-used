FROM jrei/systemd-centos:7

COPY CloudHashingInspire.viper.yaml /etc/CloudHashingInspire/CloudHashingInspire.viper.yaml
COPY cloud-hashing-inspire /usr/local/bin/cloud-hashing-inspire

CMD ["cloud-hashing-inspire", "run"]
