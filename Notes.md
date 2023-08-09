Nginx default location
/etc/nginx/sites-available/default

Command to run the website in the background
> podman run -d -p 3000:3000 localhost/my-website:latest
Make sure you have created an image already

SSL
https://certbot.eff.org/instructions?ws=nginx&os=ubuntufocal

In Containerfile
> Using podman, make sure you have made docker.io as a default registry to get
> docker images from
Useful link: https://www.youtube.com/watch?v=piwcpd_hWn0&t=706s