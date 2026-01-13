package me.dio.copa.catar.domain.model

typealias TeamDomain = Team

data class Team(
    val displayName: String
)

typealias TeamFlagDomain = TeamFlag

data class TeamFlag(
    val flagUrl: String
)