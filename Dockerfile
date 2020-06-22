FROM scratch
WORKDIR /root/myweb
COPY . /root/myweb

EXPOSE 80
CMD ["./jianshu"]
