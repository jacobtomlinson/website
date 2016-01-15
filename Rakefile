task :test do
  sh "bundle exec jekyll build"
  sh "bundle exec htmlproof ./_site --only-4xx --check-html"
end

task :serve do
  sh "bundle exec jekyll serve"
end

task :purge do
  sh "curl https://www.cloudflare.com/api_json.html -d a=fpurge_ts -d tkn=${CLOUDFLARE_API_KEY} -d email=${CLOUDFLARE_EMAIL} -d z=jacobtomlinson.co.uk -d v=1"
end
