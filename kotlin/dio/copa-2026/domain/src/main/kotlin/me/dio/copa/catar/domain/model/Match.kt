package me.dio.copa.catar.domain.model

import java.time.LocalDateTime

typealias MatchDomain = Match

data class Match(
    val id: String,
    val name: String,
    val stadium: Stadium,
    val team1: Team,
    val team2: Team,
    val team1_flag: TeamFlag,
    val team2_flag: TeamFlag,
    val date: LocalDateTime,
    val notificationEnabled: Boolean = false,
)
