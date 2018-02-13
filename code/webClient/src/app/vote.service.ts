import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { catchError, map, tap } from 'rxjs/operators';

import { Vote } from './votes/vote';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};


@Injectable()
export class VoteService {


  private heroesUrl = 'http://localhost:8080';  // URL to web api

  constructor(
    private http: HttpClient) { }

  /** GET heroes from the server */
  getVotees (): Observable<Vote[]> {
    return this.http.get<Vote[]>(this.heroesUrl)
      .pipe(
        tap(heroes => this.log(`fetched heroes`)),
        catchError(this.handleError('/votes', []))
      );
  }

  /** GET hero by id. Return `undefined` when id not found */
  getVoteNo404<Data>(id: number): Observable<Vote> {
    const url = `${this.heroesUrl}/?id=${id}`;
    return this.http.get<Vote[]>(url)
      .pipe(
        map(heroes => heroes[0]), // returns a {0|1} element array
        tap(h => {
          const outcome = h ? `fetched` : `did not find`;
          this.log(`${outcome} hero id=${id}`);
        }),
        catchError(this.handleError<Vote>(`getVote id=${id}`))
      );
  }

  /** GET hero by id. Will 404 if id not found */
  getVote(id: number): Observable<Vote> {
    const url = `${this.heroesUrl}/${id}`;
    return this.http.get<Vote>(url).pipe(
      tap(_ => this.log(`fetched hero id=${id}`)),
      catchError(this.handleError<Vote>(`getVote id=${id}`))
    );
  }

  /* GET heroes whose name contains search term */
  searchVotees(term: string): Observable<Vote[]> {
    if (!term.trim()) {
      // if not search term, return empty hero array.
      return of([]);
    }
    return this.http.get<Vote[]>(`api/heroes/?name=${term}`).pipe(
      tap(_ => this.log(`found heroes matching "${term}"`)),
      catchError(this.handleError<Vote[]>('searchVotees', []))
    );
  }

  //////// Save methods //////////

  /** POST: add a new hero to the server */
  addVote (hero: Vote): Observable<Vote> {
    return this.http.post<Vote>(this.heroesUrl, hero, httpOptions).pipe(
      tap((hero: Vote) => this.log(`added hero w/ id=${hero.id}`)),
      catchError(this.handleError<Vote>('addVote'))
    );
  }

  /** DELETE: delete the hero from the server */
  deleteVote (hero: Vote | number): Observable<Vote> {
    const id = typeof hero === 'number' ? hero : hero.id;
    const url = `${this.heroesUrl}/${id}`;

    return this.http.delete<Vote>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted hero id=${id}`)),
      catchError(this.handleError<Vote>('deleteVote'))
    );
  }

  /** PUT: update the hero on the server */
  updateVote (hero: Vote): Observable<any> {
    return this.http.put(this.heroesUrl, hero, httpOptions).pipe(
      tap(_ => this.log(`updated hero id=${hero.id}`)),
      catchError(this.handleError<any>('updateVote'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T> (operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  /** Log a VoteService message with the MessageService */
  private log(message: string) {
    this.messageService.add('VoteService: ' + message);
  }

}
