import React, { useEffect, useState } from "react";
import Header from "../../component/header";
import Sidebar from "../../component/sidebar";
import Footer from "../../component/footer";
import Swal from "sweetalert2";
import ReactPaginate from "react-paginate";
// Mendefinisikan komponen AllRoster
function AllRoster() {
  // State untuk menyimpan data pengajar yang telah difilter
  const [filteredPengajar, setFilteredPengajar] = useState([]);

  // State untuk menyimpan data mata pelajaran
  const [mataPelajaran, setMataPelajaran] = useState([]);

  // State untuk menyimpan data roster
  const [rosters, setRosters] = useState([]);

  // State untuk menyimpan data kelas
  const [kelas, setKelas] = useState([]);

  // State untuk menyimpan hari yang dipilih
  const [selectedDay, setSelectedDay] = useState("");

  // State untuk menyimpan kata kunci pencarian
  const [searchTerm, setSearchTerm] = useState("");

  // State untuk menyimpan kelas yang dipilih
  const [selectedKelas, setSelectedKelas] = useState(null);

  // Mendapatkan URL dasar dari file konfigurasi
  const { URL_ROSTER } = require("../../configapi");

  // State untuk menyimpan halaman saat ini
  const [currentPage, setCurrentPage] = useState(0);

  // Konstanta untuk jumlah item per halaman
  const PER_PAGE = 8;

  // Fetch data kelas dari API ketika komponen dimuat
  useEffect(() => {
    fetch(`${URL_ROSTER}api/kelas`)
      .then((response) => response.json())
      .then((data) => {
        if (Array.isArray(data)) {
          setKelas(data);
          handleAllClassesClick();
        }
      });
  }, []);

  // Fungsi untuk mengubah halaman saat ini
  const handlePageClick = ({ selected: selectedPage }) => {
    setCurrentPage(selectedPage);
  };

  // Fungsi untuk mengubah pengajar yang difilter berdasarkan mata pelajaran yang dipilih
  const handleMataPelajaranChange = (event) => {
    const selectedMataPelajaran = event.target.value;
    const filtered = pengajar.filter(
      (pengajarItem) => pengajarItem.Bidang === selectedMataPelajaran
    );
    setFilteredPengajar(filtered);
  };

  // Menghitung offset berdasarkan halaman saat ini
  const offset = currentPage * PER_PAGE;

  // Mendapatkan data roster untuk halaman saat ini
  const currentPageData = rosters.filter(
    (roster) =>
      (!selectedDay || roster.hari === selectedDay) &&
      (roster.mata_pelajaran.toLowerCase().includes(searchTerm.toLowerCase()) ||
        roster.pengajar.toLowerCase().includes(searchTerm.toLowerCase()))
  );

  // Menghitung jumlah halaman
  const pageCount = Math.ceil(rosters.length / PER_PAGE);
  const [isModalOpen, setIsModalOpen] = useState(false);

  // Fungsi untuk mengubah kata kunci pencarian
  const handleSearchChange = (event) => {
    setSearchTerm(event.target.value);
  };

  // Fetch semua data roster dari API ketika komponen dimuat
  useEffect(() => {
    handleAllClassesClick();
  }, []);

  // Fungsi untuk menutup modal
  const closeModal = () => {
    setIsModalOpen(false);
  };
  // Fungsi untuk fetch semua data roster dari API
  const handleAllClassesClick = () => {
    fetch(`${URL_ROSTER}api/roster`)
      .then((response) => response.json())

      .then((rosters) => setRosters(rosters))
      .catch((error) => console.error(error));
  };

  // State untuk menyimpan data pengajar
  const [pengajar, setPengajar] = useState([]);

  // Fungsi untuk mengubah hari yang dipilih
  const handleDayChange = (event) => {
    setSelectedDay(event.target.value);
  };

  // Fungsi untuk menghapus data roster
  const handleDelete = (id) => {
    Swal.fire({
      title: "Apakah Anda yakin?",
      text: "Data yang dihapus tidak dapat dikembalikan!",
      icon: "warning",
      showCancelButton: true,
      confirmButtonColor: "#3085d6",
      cancelButtonColor: "#d33",
      confirmButtonText: "Ya, hapus!",
    }).then((result) => {
      if (result.isConfirmed) {
        fetch(`${URL_ROSTER}api/roster/${id}`, {
          method: "DELETE",
        })
          .then((response) => response.json())
          .then((data) => {
            Swal.fire({
              title: "Berhasil!",
              text: "Roster berhasil dihapus",
              icon: "success",
              confirmButtonText: "OK",
            }).then(() => {
              handleAllClassesClick();
            });
          })
          .catch((error) => {
            console.error("Error:", error);
            Swal.fire({
              title: "Gagal!",
              text: "Roster gagal dihapus",
              icon: "error",
              confirmButtonText: "OK",
            });
          });
      }
    });
  };

  // Fungsi untuk menangani submit form
  const handleFormSubmit = (event) => {
    event.preventDefault();
    const form = document.forms["event-form"];
    const mata_pelajaran = form["mata_pelajaran"].value;
    const kelas = form["kelas"].value;
    const pengajar = form["pengajar"].value;
    const hari = form["hari"].value;
    const waktu_mulai = form["waktu_mulai"].value;
    const waktu_selesai = form["waktu_selesai"].value;
    const data = {
      mata_pelajaran,
      kelas,
      pengajar,
      hari,
      waktu_mulai,
      waktu_selesai,
    };
    fetch(`${URL_ROSTER}api/roster`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    })
      .then((response) => response.json())
      .then((data) => {
        Swal.fire({
          title: "Berhasil!",
          text: "Roster berhasil ditambahkan",
          icon: "success",
          confirmButtonText: "OK",
        }).then(() => {
          closeModal(); // Menutup modal setelah user mengklik OK
          handleAllClassesClick();
        });
      })
      .catch((error) => {
        console.error("Error:", error);
        Swal.fire({
          title: "Gagal!",
          text: "Roster gagal ditambahkan",
          icon: "error",
          confirmButtonText: "OK",
        });
      });
  };
  return (
    <div id="layout-wrapper">
      <div
        className="modal fade"
        id="event-modal"
        data-bs-backdrop="static"
        data-bs-keyboard="false"
        tabIndex="-1"
        style={{ display: "none" }}
        aria-hidden="true"
      >
        <div className="modal-dialog modal-dialog-centered">
          <div className="modal-content">
            <div className="modal-header py-3 px-4">
              <h5 className="modal-title" id="modal-title">
                Tambahkan Roster
              </h5>
              <button
                type="button"
                className="btn-close"
                data-bs-dismiss="modal"
                aria-label="Close"
              ></button>
            </div>
            <div className="modal-body p-4">
              <form
                className="needs-validation"
                name="event-form"
                id="form-event"
                noValidate=""
                onSubmit={handleFormSubmit}
              >
                <div className="row">
                  <div className="col-12">
                    <div className="mb-3">
                      <label className="form-label">Mata Pelajaran</label>
                      <input
                        className="form-control"
                        placeholder="Masukkan Mata Pelajaran"
                        type="text"
                        name="mata_pelajaran"
                        id="mata_pelajaran"
                        required=""
                      />
                    </div>
                    <div className="mb-3">
                      <label className="form-label">Kelas</label>
                      <input
                        className="form-control"
                        placeholder="Masukkan Kelas"
                        type="text"
                        name="kelas"
                        id="kelas"
                        required=""
                      />
                    </div>
                    <div className="mb-3">
                      <label className="form-label">Guru Pengajar</label>
                      <input
                        className="form-control"
                        placeholder="Masukkan Guru Pengajar"
                        type="text"
                        name="pengajar"
                        id="pengajar"
                        required=""
                      />
                    </div>
                    <div className="mb-3">
                      <label className="form-label">Hari</label>
                      <select
                        className="form-control"
                        name="hari"
                        id="hari"
                        required=""
                      >
                        <option value="Senin">Senin</option>
                        <option value="Selasa">Selasa</option>
                        <option value="Rabu">Rabu</option>
                        <option value="Kamis">Kamis</option>
                        <option value="Jumat">Jumat</option>
                        <option value="Sabtu">Sabtu</option>
                      </select>
                    </div>
                    <div className="mb-3">
                      <label className="form-label">Waktu Mulai</label>
                      <input
                        className="form-control"
                        placeholder="Masukkan Waktu Mulai"
                        type="time"
                        name="waktu_mulai"
                        id="waktu_mulai"
                        required=""
                      />
                    </div>
                    <div className="mb-3">
                      <label className="form-label">Waktu Selesai</label>
                      <input
                        className="form-control"
                        placeholder="Masukkan Waktu Selesai"
                        type="time"
                        name="waktu_selesai"
                        id="waktu_selesai"
                        required=""
                      />
                    </div>
                  </div>
                </div>
                <div className="row mt-2">
                  <div className="col-12 text-end">
                    <button
                      type="button"
                      className="btn btn-light me-1"
                      data-bs-dismiss="modal"
                    >
                      Close
                    </button>
                    <button
                      type="submit"
                      className="btn btn-success"
                      id="btn-save-event"
                    >
                      Save
                    </button>
                  </div>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
      <Header />
      <Sidebar />
      <div className="main-content">
        <div className="page-content">
          <div className="container-fluid">
            <div className="row">
              <div className="col-12">
                <div className="page-title-box d-sm-flex a  lign-items-center justify-content-between">
                  <div className="page-title-right">
                    <ol className="breadcrumb m-0"></ol>
                  </div>
                </div>
              </div>
            </div>

            <div className="row mb-3">
              <div className="col-xl-12">
                <div className="card mb-0">
                  <div class="card-body">
                    {selectedKelas && (
                      <>
                        <h4 class="card-title"> {selectedKelas.kelas}</h4>
                      </>
                    )}
                    <div class="row">
                      <div class="col-sm-12 col-md-9">
                        <div
                          class="dataTables_length"
                          id="complex-header-datatable_length"
                          style={{ marginLeft: "2em" }} // Add this line
                        >
                          <label>
                            Filter by Day:
                            <select
                              name="complex-header-datatable_length"
                              aria-controls="complex-header-datatable"
                              className="custom-select custom-select-sm form-control form-control-sm form-select form-select-sm"
                              style={{ width: "11em", marginTop: "0.5em" }}
                              onChange={handleDayChange} // Tambahkan ini
                            >
                              <option value="">Semua Hari </option>
                              <option value="Senin">Senin</option>
                              <option value="Selasa">Selasa</option>
                              <option value="Rabu">Rabu</option>
                              <option value="Kamis">Kamis</option>
                              <option value="Jumat">Jumat</option>
                              <option value="Sabtu">Sabtu</option>
                            </select>
                          </label>
                        </div>
                      </div>
                      <div class="col-sm-3 col-md-1">
                        <button
                          type="button"
                          class="btn font-16 btn-info waves-effect waves-light w-100"
                          id="btn-new-event"
                          data-bs-toggle="modal"
                          data-bs-target="#event-modal"
                        >
                          Add
                        </button>
                      </div>
                      <div class="col-sm-12 col-md-2">
                        <div
                          id="complex-header-datatable_filter"
                          class="dataTables_filter"
                        >
                          <label>
                            Search:
                            <input
                              type="search"
                              style={{ width: "140" }}
                              class="form-control form-control-sm"
                              placeholder="Cari disini "
                              aria-controls="complex-header-datatable"
                              onChange={handleSearchChange}
                            />
                          </label>
                        </div>
                      </div>
                    </div>
                    <div class="table-responsive">
                      <table className="table table-striped mb-0">
                        <thead>
                          <tr>
                            <th>#</th>
                            <th>Mata Pelajaran</th>
                            <th>Kelas</th>
                            <th>Guru Pengajar</th>
                            <th>Hari</th>
                            <th>Waktu Mulai</th>
                            <th>Waktu Selesai</th>
                            <th>Action</th>
                          </tr>
                        </thead>
                        <tbody>
                          {rosters
                            .filter(
                              (roster) =>
                                (!selectedDay || roster.hari === selectedDay) &&
                                (roster.mata_pelajaran
                                  .toLowerCase()
                                  .includes(searchTerm.toLowerCase()) ||
                                  roster.pengajar
                                    .toLowerCase()
                                    .includes(searchTerm.toLowerCase()))
                            )
                            .slice(offset, offset + PER_PAGE)
                            .map((roster, index) => {
                              return (
                                <tr key={index}>
                                  <td>{offset + index + 1}</td>{" "}
                                  {/* Updated this line */}
                                  <td>{roster.mata_pelajaran}</td>
                                  <td>{roster.kelas}</td>
                                  <td>{roster.pengajar} </td>
                                  <td>{roster.hari}</td>
                                  <td>{roster.waktu_mulai}</td>
                                  <td>{roster.waktu_selesai}</td>
                                  <td>
                                    <button
                                      className="btn btn-danger btn-sm"
                                      onClick={() => handleDelete(roster.id)}
                                    >
                                      Delete
                                    </button>
                                  </td>
                                </tr>
                              );
                            })}
                        </tbody>
                      </table>
                    </div>
                    <div className="row">
                      <div className="col-sm-12 col-md-5 mt-4">
                        <div
                          className="dataTables_info"
                          id="datatable-buttons_info"
                          role="status"
                          aria-live="polite"
                        >
                          Showing {currentPage * PER_PAGE + 1} to{" "}
                          {Math.min(
                            (currentPage + 1) * PER_PAGE,
                            rosters.length
                          )}{" "}
                          of {rosters.length} entries
                        </div>
                      </div>
                      <div className="col-sm-12 col-md-7">
                        <div
                          className="dataTables_paginate paging_simple_numbers"
                          id="datatable-buttons_paginate"
                        >
                          <ReactPaginate
                            previousLabel={"previous"}
                            nextLabel={"next"}
                            breakLabel={"..."}
                            breakClassName={"break-me"}
                            pageCount={pageCount}
                            marginPagesDisplayed={2}
                            pageRangeDisplayed={8}
                            onPageChange={handlePageClick}
                            containerClassName={"pagination"}
                            subContainerClassName={"pages pagination"}
                            activeClassName={"active"}
                          />
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div style={{ clear: "both" }}></div>
          </div>
        </div>

        <Footer />
      </div>
    </div>
  );
}
export default AllRoster;
