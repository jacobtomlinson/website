task :test do
  sh "bundle exec jekyll build"
  sh "bundle exec htmlproof ./_site --only-4xx --check-favicon --check-html"
end

task :serve do
  sh "bundle exec jekyll serve"
end
