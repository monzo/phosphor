FROM busybox

ADD dist/docker/bin/ /phosphor_bin/
RUN cd /    && ln -s /phosphor_bin/* . \
 && cd /bin && ln -s /phosphor_bin/* .

EXPOSE 7750 7760 7760/udp
