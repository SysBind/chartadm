# bundle

bundle "Wordpress" {

  release "database" {      
    chart "mariadb" {
      repo "stable" {
	url = "https://kubernetes-charts.storage.googleapis.com"      
      }
      version = "10.5"
    }
  }
  
  release "wordpress" {
    chart "wordpress" {
      repo "stable" {
	url = "https://kubernetes-charts.storage.googleapis.com"
      }
      version = "10"
    }
  }
}
