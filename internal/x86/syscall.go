package x86

import linux "github.com/wnxd/microdbg-linux"

const (
	NR_restart_syscall        = 0
	NR_exit                   = 1
	NR_fork                   = 2
	NR_read                   = 3
	NR_write                  = 4
	NR_open                   = 5
	NR_close                  = 6
	NR_waitpid                = 7
	NR_creat                  = 8
	NR_link                   = 9
	NR_unlink                 = 10
	NR_execve                 = 11
	NR_chdir                  = 12
	NR_time                   = 13
	NR_mknod                  = 14
	NR_chmod                  = 15
	NR_lchown                 = 16
	NR_break                  = 17
	NR_oldstat                = 18
	NR_lseek                  = 19
	NR_getpid                 = 20
	NR_mount                  = 21
	NR_umount                 = 22
	NR_setuid                 = 23
	NR_getuid                 = 24
	NR_stime                  = 25
	NR_ptrace                 = 26
	NR_alarm                  = 27
	NR_oldfstat               = 28
	NR_pause                  = 29
	NR_utime                  = 30
	NR_stty                   = 31
	NR_gtty                   = 32
	NR_access                 = 33
	NR_nice                   = 34
	NR_ftime                  = 35
	NR_sync                   = 36
	NR_kill                   = 37
	NR_rename                 = 38
	NR_mkdir                  = 39
	NR_rmdir                  = 40
	NR_dup                    = 41
	NR_pipe                   = 42
	NR_times                  = 43
	NR_prof                   = 44
	NR_brk                    = 45
	NR_setgid                 = 46
	NR_getgid                 = 47
	NR_signal                 = 48
	NR_geteuid                = 49
	NR_getegid                = 50
	NR_acct                   = 51
	NR_umount2                = 52
	NR_lock                   = 53
	NR_ioctl                  = 54
	NR_fcntl                  = 55
	NR_mpx                    = 56
	NR_setpgid                = 57
	NR_ulimit                 = 58
	NR_oldolduname            = 59
	NR_umask                  = 60
	NR_chroot                 = 61
	NR_ustat                  = 62
	NR_dup2                   = 63
	NR_getppid                = 64
	NR_getpgrp                = 65
	NR_setsid                 = 66
	NR_sigaction              = 67
	NR_sgetmask               = 68
	NR_ssetmask               = 69
	NR_setreuid               = 70
	NR_setregid               = 71
	NR_sigsuspend             = 72
	NR_sigpending             = 73
	NR_sethostname            = 74
	NR_setrlimit              = 75
	NR_getrlimit              = 76
	NR_getrusage              = 77
	NR_gettimeofday           = 78
	NR_settimeofday           = 79
	NR_getgroups              = 80
	NR_setgroups              = 81
	NR_select                 = 82
	NR_symlink                = 83
	NR_oldlstat               = 84
	NR_readlink               = 85
	NR_uselib                 = 86
	NR_swapon                 = 87
	NR_reboot                 = 88
	NR_readdir                = 89
	NR_mmap                   = 90
	NR_munmap                 = 91
	NR_truncate               = 92
	NR_ftruncate              = 93
	NR_fchmod                 = 94
	NR_fchown                 = 95
	NR_getpriority            = 96
	NR_setpriority            = 97
	NR_profil                 = 98
	NR_statfs                 = 99
	NR_fstatfs                = 100
	NR_ioperm                 = 101
	NR_socketcall             = 102
	NR_syslog                 = 103
	NR_setitimer              = 104
	NR_getitimer              = 105
	NR_stat                   = 106
	NR_lstat                  = 107
	NR_fstat                  = 108
	NR_olduname               = 109
	NR_iopl                   = 110
	NR_vhangup                = 111
	NR_idle                   = 112
	NR_vm86old                = 113
	NR_wait4                  = 114
	NR_swapoff                = 115
	NR_sysinfo                = 116
	NR_ipc                    = 117
	NR_fsync                  = 118
	NR_sigreturn              = 119
	NR_clone                  = 120
	NR_setdomainname          = 121
	NR_uname                  = 122
	NR_modify_ldt             = 123
	NR_adjtimex               = 124
	NR_mprotect               = 125
	NR_sigprocmask            = 126
	NR_create_module          = 127
	NR_init_module            = 128
	NR_delete_module          = 129
	NR_get_kernel_syms        = 130
	NR_quotactl               = 131
	NR_getpgid                = 132
	NR_fchdir                 = 133
	NR_bdflush                = 134
	NR_sysfs                  = 135
	NR_personality            = 136
	NR_afs_syscall            = 137
	NR_setfsuid               = 138
	NR_setfsgid               = 139
	NR__llseek                = 140
	NR_getdents               = 141
	NR__newselect             = 142
	NR_flock                  = 143
	NR_msync                  = 144
	NR_readv                  = 145
	NR_writev                 = 146
	NR_getsid                 = 147
	NR_fdatasync              = 148
	NR__sysctl                = 149
	NR_mlock                  = 150
	NR_munlock                = 151
	NR_mlockall               = 152
	NR_munlockall             = 153
	NR_sched_setparam         = 154
	NR_sched_getparam         = 155
	NR_sched_setscheduler     = 156
	NR_sched_getscheduler     = 157
	NR_sched_yield            = 158
	NR_sched_get_priority_max = 159
	NR_sched_get_priority_min = 160
	NR_sched_rr_get_interval  = 161
	NR_nanosleep              = 162
	NR_mremap                 = 163
	NR_setresuid              = 164
	NR_getresuid              = 165
	NR_vm86                   = 166
	NR_query_module           = 167
	NR_poll                   = 168
	NR_nfsservctl             = 169
	NR_setresgid              = 170
	NR_getresgid              = 171
	NR_prctl                  = 172
	NR_rt_sigreturn           = 173
	NR_rt_sigaction           = 174
	NR_rt_sigprocmask         = 175
	NR_rt_sigpending          = 176
	NR_rt_sigtimedwait        = 177
	NR_rt_sigqueueinfo        = 178
	NR_rt_sigsuspend          = 179
	NR_pread64                = 180
	NR_pwrite64               = 181
	NR_chown                  = 182
	NR_getcwd                 = 183
	NR_capget                 = 184
	NR_capset                 = 185
	NR_sigaltstack            = 186
	NR_sendfile               = 187
	NR_getpmsg                = 188
	NR_putpmsg                = 189
	NR_vfork                  = 190
	NR_ugetrlimit             = 191
	NR_mmap2                  = 192
	NR_truncate64             = 193
	NR_ftruncate64            = 194
	NR_stat64                 = 195
	NR_lstat64                = 196
	NR_fstat64                = 197
	NR_lchown32               = 198
	NR_getuid32               = 199
	NR_getgid32               = 200
	NR_geteuid32              = 201
	NR_getegid32              = 202
	NR_setreuid32             = 203
	NR_setregid32             = 204
	NR_getgroups32            = 205
	NR_setgroups32            = 206
	NR_fchown32               = 207
	NR_setresuid32            = 208
	NR_getresuid32            = 209
	NR_setresgid32            = 210
	NR_getresgid32            = 211
	NR_chown32                = 212
	NR_setuid32               = 213
	NR_setgid32               = 214
	NR_setfsuid32             = 215
	NR_setfsgid32             = 216
	NR_pivot_root             = 217
	NR_mincore                = 218
	NR_madvise                = 219
	NR_getdents64             = 220
	NR_fcntl64                = 221
	NR_gettid                 = 224
	NR_readahead              = 225
	NR_setxattr               = 226
	NR_lsetxattr              = 227
	NR_fsetxattr              = 228
	NR_getxattr               = 229
	NR_lgetxattr              = 230
	NR_fgetxattr              = 231
	NR_listxattr              = 232
	NR_llistxattr             = 233
	NR_flistxattr             = 234
	NR_removexattr            = 235
	NR_lremovexattr           = 236
	NR_fremovexattr           = 237
	NR_tkill                  = 238
	NR_sendfile64             = 239
	NR_futex                  = 240
	NR_sched_setaffinity      = 241
	NR_sched_getaffinity      = 242
	NR_set_thread_area        = 243
	NR_get_thread_area        = 244
	NR_io_setup               = 245
	NR_io_destroy             = 246
	NR_io_getevents           = 247
	NR_io_submit              = 248
	NR_io_cancel              = 249
	NR_fadvise64              = 250
	NR_exit_group             = 252
	NR_lookup_dcookie         = 253
	NR_epoll_create           = 254
	NR_epoll_ctl              = 255
	NR_epoll_wait             = 256
	NR_remap_file_pages       = 257
	NR_set_tid_address        = 258
	NR_timer_create           = 259
	NR_timer_settime          = 260
	NR_timer_gettime          = 261
	NR_timer_getoverrun       = 262
	NR_timer_delete           = 263
	NR_clock_settime          = 264
	NR_clock_gettime          = 265
	NR_clock_getres           = 266
	NR_clock_nanosleep        = 267
	NR_statfs64               = 268
	NR_fstatfs64              = 269
	NR_tgkill                 = 270
	NR_utimes                 = 271
	NR_fadvise64_64           = 272
	NR_vserver                = 273
	NR_mbind                  = 274
	NR_get_mempolicy          = 275
	NR_set_mempolicy          = 276
	NR_mq_open                = 277
	NR_mq_unlink              = 278
	NR_mq_timedsend           = 279
	NR_mq_timedreceive        = 280
	NR_mq_notify              = 281
	NR_mq_getsetattr          = 282
	NR_kexec_load             = 283
	NR_waitid                 = 284
	NR_add_key                = 286
	NR_request_key            = 287
	NR_keyctl                 = 288
	NR_ioprio_set             = 289
	NR_ioprio_get             = 290
	NR_inotify_init           = 291
	NR_inotify_add_watch      = 292
	NR_inotify_rm_watch       = 293
	NR_migrate_pages          = 294
	NR_openat                 = 295
	NR_mkdirat                = 296
	NR_mknodat                = 297
	NR_fchownat               = 298
	NR_futimesat              = 299
	NR_fstatat64              = 300
	NR_unlinkat               = 301
	NR_renameat               = 302
	NR_linkat                 = 303
	NR_symlinkat              = 304
	NR_readlinkat             = 305
	NR_fchmodat               = 306
	NR_faccessat              = 307
	NR_pselect6               = 308
	NR_ppoll                  = 309
	NR_unshare                = 310
	NR_set_robust_list        = 311
	NR_get_robust_list        = 312
	NR_splice                 = 313
	NR_sync_file_range        = 314
	NR_tee                    = 315
	NR_vmsplice               = 316
	NR_move_pages             = 317
	NR_getcpu                 = 318
	NR_epoll_pwait            = 319
	NR_utimensat              = 320
	NR_signalfd               = 321
	NR_timerfd_create         = 322
	NR_eventfd                = 323
	NR_fallocate              = 324
	NR_timerfd_settime        = 325
	NR_timerfd_gettime        = 326
	NR_signalfd4              = 327
	NR_eventfd2               = 328
	NR_epoll_create1          = 329
	NR_dup3                   = 330
	NR_pipe2                  = 331
	NR_inotify_init1          = 332
	NR_preadv                 = 333
	NR_pwritev                = 334
	NR_rt_tgsigqueueinfo      = 335
	NR_perf_event_open        = 336
	NR_recvmmsg               = 337
	NR_fanotify_init          = 338
	NR_fanotify_mark          = 339
	NR_prlimit64              = 340
	NR_name_to_handle_at      = 341
	NR_open_by_handle_at      = 342
	NR_clock_adjtime          = 343
	NR_syncfs                 = 344
	NR_sendmmsg               = 345
	NR_setns                  = 346
	NR_process_vm_readv       = 347
	NR_process_vm_writev      = 348
	NR_kcmp                   = 349
	NR_finit_module           = 350
	NR_sched_setattr          = 351
	NR_sched_getattr          = 352
	NR_renameat2              = 353
	NR_seccomp                = 354
	NR_getrandom              = 355
	NR_memfd_create           = 356
	NR_bpf                    = 357
	NR_execveat               = 358
	NR_socket                 = 359
	NR_socketpair             = 360
	NR_bind                   = 361
	NR_connect                = 362
	NR_listen                 = 363
	NR_accept4                = 364
	NR_getsockopt             = 365
	NR_setsockopt             = 366
	NR_getsockname            = 367
	NR_getpeername            = 368
	NR_sendto                 = 369
	NR_sendmsg                = 370
	NR_recvfrom               = 371
	NR_recvmsg                = 372
	NR_shutdown               = 373
	NR_userfaultfd            = 374
	NR_membarrier             = 375
	NR_mlock2                 = 376
	NR_copy_file_range        = 377
	NR_preadv2                = 378
	NR_pwritev2               = 379
	NR_pkey_mprotect          = 380
	NR_pkey_alloc             = 381
	NR_pkey_free              = 382
	NR_statx                  = 383
	NR_arch_prctl             = 384
	NR_io_pgetevents          = 385
	NR_rseq                   = 386
)

var (
	NRMap = map[uint64]linux.NR{
		NR_io_setup:               linux.NR_io_setup,
		NR_io_destroy:             linux.NR_io_destroy,
		NR_io_submit:              linux.NR_io_submit,
		NR_io_cancel:              linux.NR_io_cancel,
		NR_io_getevents:           linux.NR_io_getevents,
		NR_setxattr:               linux.NR_setxattr,
		NR_lsetxattr:              linux.NR_lsetxattr,
		NR_fsetxattr:              linux.NR_fsetxattr,
		NR_getxattr:               linux.NR_getxattr,
		NR_lgetxattr:              linux.NR_lgetxattr,
		NR_fgetxattr:              linux.NR_fgetxattr,
		NR_listxattr:              linux.NR_listxattr,
		NR_llistxattr:             linux.NR_llistxattr,
		NR_flistxattr:             linux.NR_flistxattr,
		NR_removexattr:            linux.NR_removexattr,
		NR_lremovexattr:           linux.NR_lremovexattr,
		NR_fremovexattr:           linux.NR_fremovexattr,
		NR_getcwd:                 linux.NR_getcwd,
		NR_lookup_dcookie:         linux.NR_lookup_dcookie,
		NR_eventfd2:               linux.NR_eventfd2,
		NR_epoll_create1:          linux.NR_epoll_create1,
		NR_epoll_ctl:              linux.NR_epoll_ctl,
		NR_epoll_pwait:            linux.NR_epoll_pwait,
		NR_dup:                    linux.NR_dup,
		NR_dup3:                   linux.NR_dup3,
		NR_fcntl:                  linux.NR_fcntl,
		NR_inotify_init1:          linux.NR_inotify_init1,
		NR_inotify_add_watch:      linux.NR_inotify_add_watch,
		NR_inotify_rm_watch:       linux.NR_inotify_rm_watch,
		NR_ioctl:                  linux.NR_ioctl,
		NR_ioprio_set:             linux.NR_ioprio_set,
		NR_ioprio_get:             linux.NR_ioprio_get,
		NR_flock:                  linux.NR_flock,
		NR_mknodat:                linux.NR_mknodat,
		NR_mkdirat:                linux.NR_mkdirat,
		NR_unlinkat:               linux.NR_unlinkat,
		NR_symlinkat:              linux.NR_symlinkat,
		NR_linkat:                 linux.NR_linkat,
		NR_renameat:               linux.NR_renameat,
		NR_umount2:                linux.NR_umount2,
		NR_mount:                  linux.NR_mount,
		NR_pivot_root:             linux.NR_pivot_root,
		NR_nfsservctl:             linux.NR_nfsservctl,
		NR_statfs:                 linux.NR_statfs,
		NR_fstatfs:                linux.NR_fstatfs,
		NR_truncate:               linux.NR_truncate,
		NR_ftruncate:              linux.NR_ftruncate,
		NR_fallocate:              linux.NR_fallocate,
		NR_faccessat:              linux.NR_faccessat,
		NR_chdir:                  linux.NR_chdir,
		NR_fchdir:                 linux.NR_fchdir,
		NR_chroot:                 linux.NR_chroot,
		NR_fchmod:                 linux.NR_fchmod,
		NR_fchmodat:               linux.NR_fchmodat,
		NR_fchownat:               linux.NR_fchownat,
		NR_fchown:                 linux.NR_fchown,
		NR_open:                   linux.NR_open,
		NR_openat:                 linux.NR_openat,
		NR_close:                  linux.NR_close,
		NR_vhangup:                linux.NR_vhangup,
		NR_pipe2:                  linux.NR_pipe2,
		NR_quotactl:               linux.NR_quotactl,
		NR_getdents64:             linux.NR_getdents64,
		NR_lseek:                  linux.NR_lseek,
		NR_read:                   linux.NR_read,
		NR_write:                  linux.NR_write,
		NR_readv:                  linux.NR_readv,
		NR_writev:                 linux.NR_writev,
		NR_pread64:                linux.NR_pread64,
		NR_pwrite64:               linux.NR_pwrite64,
		NR_preadv:                 linux.NR_preadv,
		NR_pwritev:                linux.NR_pwritev,
		NR_sendfile:               linux.NR_sendfile,
		NR_pselect6:               linux.NR_pselect6,
		NR_ppoll:                  linux.NR_ppoll,
		NR_signalfd4:              linux.NR_signalfd4,
		NR_vmsplice:               linux.NR_vmsplice,
		NR_splice:                 linux.NR_splice,
		NR_tee:                    linux.NR_tee,
		NR_readlinkat:             linux.NR_readlinkat,
		NR_fstatat64:              linux.NR_fstatat64,
		NR_fstat64:                linux.NR_fstat64,
		NR_sync:                   linux.NR_sync,
		NR_fsync:                  linux.NR_fsync,
		NR_fdatasync:              linux.NR_fdatasync,
		NR_sync_file_range:        linux.NR_sync_file_range,
		NR_timerfd_create:         linux.NR_timerfd_create,
		NR_timerfd_settime:        linux.NR_timerfd_settime,
		NR_timerfd_gettime:        linux.NR_timerfd_gettime,
		NR_utimensat:              linux.NR_utimensat,
		NR_acct:                   linux.NR_acct,
		NR_capget:                 linux.NR_capget,
		NR_capset:                 linux.NR_capset,
		NR_personality:            linux.NR_personality,
		NR_exit:                   linux.NR_exit,
		NR_exit_group:             linux.NR_exit_group,
		NR_waitid:                 linux.NR_waitid,
		NR_set_tid_address:        linux.NR_set_tid_address,
		NR_unshare:                linux.NR_unshare,
		NR_futex:                  linux.NR_futex,
		NR_set_robust_list:        linux.NR_set_robust_list,
		NR_get_robust_list:        linux.NR_get_robust_list,
		NR_nanosleep:              linux.NR_nanosleep,
		NR_getitimer:              linux.NR_getitimer,
		NR_setitimer:              linux.NR_setitimer,
		NR_kexec_load:             linux.NR_kexec_load,
		NR_init_module:            linux.NR_init_module,
		NR_delete_module:          linux.NR_delete_module,
		NR_timer_create:           linux.NR_timer_create,
		NR_timer_gettime:          linux.NR_timer_gettime,
		NR_timer_getoverrun:       linux.NR_timer_getoverrun,
		NR_timer_settime:          linux.NR_timer_settime,
		NR_timer_delete:           linux.NR_timer_delete,
		NR_clock_settime:          linux.NR_clock_settime,
		NR_clock_gettime:          linux.NR_clock_gettime,
		NR_clock_getres:           linux.NR_clock_getres,
		NR_clock_nanosleep:        linux.NR_clock_nanosleep,
		NR_syslog:                 linux.NR_syslog,
		NR_ptrace:                 linux.NR_ptrace,
		NR_sched_setparam:         linux.NR_sched_setparam,
		NR_sched_setscheduler:     linux.NR_sched_setscheduler,
		NR_sched_getscheduler:     linux.NR_sched_getscheduler,
		NR_sched_getparam:         linux.NR_sched_getparam,
		NR_sched_setaffinity:      linux.NR_sched_setaffinity,
		NR_sched_getaffinity:      linux.NR_sched_getaffinity,
		NR_sched_yield:            linux.NR_sched_yield,
		NR_sched_get_priority_max: linux.NR_sched_get_priority_max,
		NR_sched_get_priority_min: linux.NR_sched_get_priority_min,
		NR_sched_rr_get_interval:  linux.NR_sched_rr_get_interval,
		NR_restart_syscall:        linux.NR_restart_syscall,
		NR_kill:                   linux.NR_kill,
		NR_tkill:                  linux.NR_tkill,
		NR_tgkill:                 linux.NR_tgkill,
		NR_sigaltstack:            linux.NR_sigaltstack,
		NR_rt_sigsuspend:          linux.NR_rt_sigsuspend,
		NR_rt_sigaction:           linux.NR_rt_sigaction,
		NR_rt_sigprocmask:         linux.NR_rt_sigprocmask,
		NR_rt_sigpending:          linux.NR_rt_sigpending,
		NR_rt_sigtimedwait:        linux.NR_rt_sigtimedwait,
		NR_rt_sigqueueinfo:        linux.NR_rt_sigqueueinfo,
		NR_rt_sigreturn:           linux.NR_rt_sigreturn,
		NR_setpriority:            linux.NR_setpriority,
		NR_getpriority:            linux.NR_getpriority,
		NR_reboot:                 linux.NR_reboot,
		NR_setregid:               linux.NR_setregid,
		NR_setgid:                 linux.NR_setgid,
		NR_setreuid:               linux.NR_setreuid,
		NR_setuid:                 linux.NR_setuid,
		NR_setresuid:              linux.NR_setresuid,
		NR_getresuid:              linux.NR_getresuid,
		NR_setresgid:              linux.NR_setresgid,
		NR_getresgid:              linux.NR_getresgid,
		NR_setfsuid:               linux.NR_setfsuid,
		NR_setfsgid:               linux.NR_setfsgid,
		NR_times:                  linux.NR_times,
		NR_setpgid:                linux.NR_setpgid,
		NR_getpgid:                linux.NR_getpgid,
		NR_getsid:                 linux.NR_getsid,
		NR_setsid:                 linux.NR_setsid,
		NR_getgroups:              linux.NR_getgroups,
		NR_setgroups:              linux.NR_setgroups,
		NR_uname:                  linux.NR_uname,
		NR_sethostname:            linux.NR_sethostname,
		NR_setdomainname:          linux.NR_setdomainname,
		NR_getrlimit:              linux.NR_getrlimit,
		NR_setrlimit:              linux.NR_setrlimit,
		NR_getrusage:              linux.NR_getrusage,
		NR_umask:                  linux.NR_umask,
		NR_prctl:                  linux.NR_prctl,
		NR_getcpu:                 linux.NR_getcpu,
		NR_gettimeofday:           linux.NR_gettimeofday,
		NR_settimeofday:           linux.NR_settimeofday,
		NR_adjtimex:               linux.NR_adjtimex,
		NR_getpid:                 linux.NR_getpid,
		NR_getppid:                linux.NR_getppid,
		NR_getuid32:               linux.NR_getuid,
		NR_geteuid32:              linux.NR_geteuid,
		NR_getgid32:               linux.NR_getgid,
		NR_getegid32:              linux.NR_getegid,
		NR_gettid:                 linux.NR_gettid,
		NR_sysinfo:                linux.NR_sysinfo,
		NR_mq_open:                linux.NR_mq_open,
		NR_mq_unlink:              linux.NR_mq_unlink,
		NR_mq_timedsend:           linux.NR_mq_timedsend,
		NR_mq_timedreceive:        linux.NR_mq_timedreceive,
		NR_mq_notify:              linux.NR_mq_notify,
		NR_mq_getsetattr:          linux.NR_mq_getsetattr,
		NR_socket:                 linux.NR_socket,
		NR_socketpair:             linux.NR_socketpair,
		NR_bind:                   linux.NR_bind,
		NR_listen:                 linux.NR_listen,
		NR_connect:                linux.NR_connect,
		NR_getsockname:            linux.NR_getsockname,
		NR_getpeername:            linux.NR_getpeername,
		NR_sendto:                 linux.NR_sendto,
		NR_recvfrom:               linux.NR_recvfrom,
		NR_setsockopt:             linux.NR_setsockopt,
		NR_getsockopt:             linux.NR_getsockopt,
		NR_shutdown:               linux.NR_shutdown,
		NR_sendmsg:                linux.NR_sendmsg,
		NR_recvmsg:                linux.NR_recvmsg,
		NR_readahead:              linux.NR_readahead,
		NR_brk:                    linux.NR_brk,
		NR_munmap:                 linux.NR_munmap,
		NR_mremap:                 linux.NR_mremap,
		NR_add_key:                linux.NR_add_key,
		NR_request_key:            linux.NR_request_key,
		NR_keyctl:                 linux.NR_keyctl,
		NR_clone:                  linux.NR_clone,
		NR_execve:                 linux.NR_execve,
		NR_mmap:                   linux.NR_mmap,
		NR_mmap2:                  linux.NR_mmap2,
		NR_fadvise64:              linux.NR_fadvise64,
		NR_swapon:                 linux.NR_swapon,
		NR_swapoff:                linux.NR_swapoff,
		NR_mprotect:               linux.NR_mprotect,
		NR_msync:                  linux.NR_msync,
		NR_mlock:                  linux.NR_mlock,
		NR_munlock:                linux.NR_munlock,
		NR_mlockall:               linux.NR_mlockall,
		NR_munlockall:             linux.NR_munlockall,
		NR_mincore:                linux.NR_mincore,
		NR_madvise:                linux.NR_madvise,
		NR_remap_file_pages:       linux.NR_remap_file_pages,
		NR_mbind:                  linux.NR_mbind,
		NR_get_mempolicy:          linux.NR_get_mempolicy,
		NR_set_mempolicy:          linux.NR_set_mempolicy,
		NR_migrate_pages:          linux.NR_migrate_pages,
		NR_move_pages:             linux.NR_move_pages,
		NR_rt_tgsigqueueinfo:      linux.NR_rt_tgsigqueueinfo,
		NR_perf_event_open:        linux.NR_perf_event_open,
		NR_accept4:                linux.NR_accept4,
		NR_recvmmsg:               linux.NR_recvmmsg,
		NR_wait4:                  linux.NR_wait4,
		NR_prlimit64:              linux.NR_prlimit64,
		NR_fanotify_init:          linux.NR_fanotify_init,
		NR_fanotify_mark:          linux.NR_fanotify_mark,
		NR_name_to_handle_at:      linux.NR_name_to_handle_at,
		NR_open_by_handle_at:      linux.NR_open_by_handle_at,
		NR_clock_adjtime:          linux.NR_clock_adjtime,
		NR_syncfs:                 linux.NR_syncfs,
		NR_setns:                  linux.NR_setns,
		NR_sendmmsg:               linux.NR_sendmmsg,
		NR_process_vm_readv:       linux.NR_process_vm_readv,
		NR_process_vm_writev:      linux.NR_process_vm_writev,
		NR_kcmp:                   linux.NR_kcmp,
		NR_finit_module:           linux.NR_finit_module,
		NR_sched_setattr:          linux.NR_sched_setattr,
		NR_sched_getattr:          linux.NR_sched_getattr,
		NR_renameat2:              linux.NR_renameat2,
		NR_seccomp:                linux.NR_seccomp,
		NR_getrandom:              linux.NR_getrandom,
		NR_memfd_create:           linux.NR_memfd_create,
		NR_bpf:                    linux.NR_bpf,
		NR_execveat:               linux.NR_execveat,
		NR_userfaultfd:            linux.NR_userfaultfd,
		NR_membarrier:             linux.NR_membarrier,
		NR_mlock2:                 linux.NR_mlock2,
		NR_copy_file_range:        linux.NR_copy_file_range,
		NR_preadv2:                linux.NR_preadv2,
		NR_pwritev2:               linux.NR_pwritev2,
		NR_pkey_mprotect:          linux.NR_pkey_mprotect,
		NR_pkey_alloc:             linux.NR_pkey_alloc,
		NR_pkey_free:              linux.NR_pkey_free,
		NR_statx:                  linux.NR_statx,
		NR_io_pgetevents:          linux.NR_io_pgetevents,
		NR_rseq:                   linux.NR_rseq,
	}
)
