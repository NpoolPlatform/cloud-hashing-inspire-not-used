FROM jrei/systemd-centos:7

COPY CloudHashingStimulate.viper.yaml /etc/CloudHashingStimulate/CloudHashingStimulate.viper.yaml
COPY cloud-hashing-stimulate /usr/local/bin/cloud-hashing-stimulate

CMD ["cloud-hashing-stimulate", "run"]
