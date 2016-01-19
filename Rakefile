task :test do
  sh "bundle exec jekyll build"
  sh "bundle exec htmlproof ./_site --only-4xx --check-html --verbose"
end

task :serve do
  sh "bundle exec jekyll serve"
end

task :purge do
  sh "curl -w \"\n\" -X DELETE \
      \"https://api.cloudflare.com/client/v4/zones/${CLOUDFLARE_ZONE_IDENTIFIER}/purge_cache\" \
      -H \"X-Auth-Email: ${CLOUDFLARE_EMAIL}\" \
      -H \"X-Auth-Key: ${CLOUDFLARE_API_KEY}\" \
      -H \"Content-Type: application/json\" \
      --data '{\"purge_everything\":true}'"
end
