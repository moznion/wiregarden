plugins {
    `java-library`
    `maven-publish`
    signing
}

repositories {
    mavenCentral()
}

dependencies {
    implementation("com.google.protobuf:protobuf-java:3.25.1")
    implementation("io.grpc:grpc-protobuf:1.60.1")
    implementation("io.grpc:grpc-stub:1.60.1")
    testImplementation("org.junit.jupiter:junit-jupiter-api:5.10.1")
    testRuntimeOnly("org.junit.jupiter:junit-jupiter-engine")
}

java {
    sourceCompatibility = JavaVersion.VERSION_1_8
    targetCompatibility = JavaVersion.VERSION_1_8
    withJavadocJar()
    withSourcesJar()
}

tasks.test {
    useJUnitPlatform()
}

publishing {
    publications {
        create<MavenPublication>("maven") {
            group = "net.moznion"
            artifactId = "wiregarden"
            version = "0.5.2"
            from(components["java"])
            pom {
                name.set("wiregarden")
                description.set("A library for wiregarden's gRPC")
                url.set("https://github.com/moznion/wiregarden")
                licenses {
                    license {
                        name.set("MIT License")
                        url.set("https://github.com/moznion/wiregarden/blob/main/LICENSE")
                    }
                }
                developers {
                    developer {
                        id.set("moznion")
                        name.set("Taiki Kawakami")
                        email.set("moznion@mail.moznion.net")
                    }
                }
                scm {
                    connection.set("scm:git:git://github.com/moznion/wiregarden.git")
                    developerConnection.set("scm:git:ssh://github.com/moznion/wiregarden.git")
                    url.set("https://github.com/moznion/wiregarden/ext/lib/java")
                }
            }
        }
    }
    repositories {
        maven {
            val releasesRepoUrl: String = "https://oss.sonatype.org/service/local/staging/deploy/maven2"
            val snapshotsRepoUrl: String = "https://oss.sonatype.org/content/repositories/snapshots"
            setUrl(uri(if ((version as String).endsWith("SNAPSHOT")) snapshotsRepoUrl else releasesRepoUrl))
            credentials {
                username = fun (): String {
                    val sonatypeUsername = findProperty("sonatypeUsername") ?: return ""
                    return sonatypeUsername as String
                }()
                password = fun (): String {
                    val sonatypePassword = findProperty("sonatypePassword") ?: return ""
                    return sonatypePassword as String
                }()
            }
        }
    }
}

signing {
    sign(publishing.publications["maven"])
}
