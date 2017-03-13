
class GCloudConfig
  attr_accessor :region, :config, :name

  def initialize( name, args = {} )
    @region = args[:region] || 'us-central'
    @config = args[:config] || 'default'

    @name = name
  end

  def url
    "http://#{name}.appspot.com/"
  end
end

camhd_app_dev = GCloudConfig.new( 'camhd-app-dev',
                        config: 'camhd-app-dev' )


namespace :gcloud do

  task :activate do
    sh *%W( gcloud config configurations activate #{camhd_app_dev.config} )
  end

  task :create => :activate do
    sh *%W( gcloud app create --region=#{camhd_app_dev.region})
  end

  task :deploy => :activate do
    sh *%w( gcloud app deploy --stop-previous-version app.yaml )
  end

  task :open do
    sh *%w( open #{camhd_app_dev.url()} )
  end

end

namespace :docker do

task :build do
  sh *%w( docker build -t go-lazycache-appengine:dev . )
end

task :run do
  sh *%w( docker run --rm --publish 8080:8080 go-lazycache-appengine:dev )
end

task :console do
  sh *%w( docker run --rm --tty --interactive --entrypoint /bin/bash --publish-all go-lazycache-appengine:dev )
end

end
