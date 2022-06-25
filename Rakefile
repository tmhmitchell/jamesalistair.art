require "tinify"
require "pathname"

task :img do
    Tinify.key = ENV["TINIFY_API_KEY"]

    pn = Pathname.new(ARGV[1])
    dest_filename = pn.dirname().join(pn.basename(".*").to_s() + "-compressed.jpg")

    tinify_source = Tinify.from_file(pn.realpath())
    tinify_source.to_file(dest_filename)

    puts "Written file to: " + dest_filename.to_s()
end

task :serve do
    sh "bundle exec jekyll serve --livereload --trace"
end