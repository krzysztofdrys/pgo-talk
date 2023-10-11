FROM krzysztofdrys/developmentmachine

RUN /usr/local/go/bin/go install golang.org/x/perf/cmd/benchstat@latest
RUN ln -s ${HOME}/go/bin/benchstat ${HOME}/bin

RUN /usr/local/go/bin/go install github.com/felixge/pprofutils/v2/cmd/pprofutils@latest
RUN ln -s ${HOME}/go/bin/pprofutils ${HOME}/bin

RUN wget 'https://github.com/sharkdp/hyperfine/releases/download/v1.17.0/hyperfine-v1.17.0-x86_64-unknown-linux-gnu.tar.gz'
RUN tar -xvf hyperfine-v1.17.0-x86_64-unknown-linux-gnu.tar.gz
RUN ln -s ${HOME}/hyperfine-v1.17.0-x86_64-unknown-linux-gnu/hyperfine ${HOME}/bin

