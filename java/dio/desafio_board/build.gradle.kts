plugins {
    id("java")
}

group = "com.taranttini"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    // driver to connection
    implementation("com.mysql:mysql-connector-j:8.3.0")

    // used to made migrations
    implementation("org.flywaydb:flyway-core:10.11.1")
    implementation("org.flywaydb:flyway-mysql:10.11.1")


    implementation("org.projectlombok:lombok:1.18.34")

    annotationProcessor("org.projectlombok:lombok:1.18.34")
}


tasks.test {
    useJUnitPlatform()
}